package client

import (
	"context"
	"encoding/json"

	"github.com/LobovVit/GophKeeper/internal/client/constants"
	"github.com/LobovVit/GophKeeper/internal/client/model"
	grpc "github.com/LobovVit/GophKeeper/internal/proto"
	"github.com/LobovVit/GophKeeper/pkg/logger"
	"github.com/LobovVit/GophKeeper/pkg/utils"
	"go.uber.org/zap"
)

// Synchronization - run synchronization
func (c Client) Synchronization(ctx context.Context, password string, token model.Token) ([][]string, [][]string, [][]string, [][]string, error) {
	logger.Log.Info("synchronization")

	dataTblText := [][]string{}
	dataTblCard := [][]string{}
	dataTblLoginPassword := [][]string{}
	dataTblBinary := [][]string{}

	created := utils.ConvertTimeToTimestamp(token.CreatedAt)
	endDate := utils.ConvertTimeToTimestamp(token.EndDateAt)
	//-----------------------------------------------
	var plaintext string
	secretKey := utils.AesKeySecureRandom([]byte(password))

	titleText := []string{constants.NameItem, constants.DescriptionItem, constants.DataItem, constants.CreatedAtItem, constants.UpdatedAtItem}
	titleCard := []string{constants.NameItem, constants.DescriptionItem, constants.PaymentSystemItem, constants.NumberItem, constants.HolderItem,
		constants.CVCItem, constants.EndDateItem, constants.CreatedAtItem, constants.UpdatedAtItem}
	titleLoginPassword := []string{constants.NameItem, constants.DescriptionItem, constants.LoginItem, constants.PasswordItem,
		constants.CreatedAtItem, constants.UpdatedAtItem}
	titleBinary := []string{constants.NameItem, constants.CreatedAtItem}
	dataTblText = append(dataTblText, titleText)
	dataTblCard = append(dataTblCard, titleCard)
	dataTblLoginPassword = append(dataTblLoginPassword, titleLoginPassword)
	dataTblBinary = append(dataTblBinary, titleBinary)

	dataTblTextPointer := &dataTblText
	dataTblCardPointer := &dataTblCard
	dataTblLoginPasswordPointer := &dataTblLoginPassword
	dataTblBinaryPointer := &dataTblBinary
	//-----------------------------------------------
	nodesTextEntity, err := c.grpc.EntityGetList(ctx,
		&grpc.GetListEntityRequest{Type: constants.Text.ToString(),
			AccessToken: &grpc.Token{Token: token.AccessToken,
				UserId: token.UserID, CreatedAt: created, EndDateAt: endDate}})
	if err != nil {
		logger.Log.Error("entity get list", zap.Error(err))
		return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
	}
	for _, node := range nodesTextEntity.Node {
		plaintext, err = utils.Decrypt(string(node.Data), secretKey)
		if err != nil {
			logger.Log.Error("decrypt", zap.Error(err))
			return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
		}
		err = AppendTextEntity(node, dataTblTextPointer, plaintext)
		if err != nil {
			logger.Log.Error("append text entity", zap.Error(err))
			return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
		}
	}

	//-----------------------------------------------

	nodesCardEntity, err := c.grpc.EntityGetList(ctx,
		&grpc.GetListEntityRequest{Type: constants.Card.ToString(),
			AccessToken: &grpc.Token{Token: token.AccessToken,
				UserId: token.UserID, CreatedAt: created, EndDateAt: endDate}})
	if err != nil {
		logger.Log.Error("entity get list", zap.Error(err))
		return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
	}
	for _, node := range nodesCardEntity.Node {
		plaintext, err = utils.Decrypt(string(node.Data), secretKey)
		if err != nil {
			logger.Log.Error("decrypt", zap.Error(err))
			return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
		}

		var card model.Card
		err = json.Unmarshal([]byte(plaintext), &card)
		if err != nil {
			logger.Log.Error("unmarshal", zap.Error(err))
			return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
		}
		err = AppendCardEntity(node, dataTblCardPointer, card)
		if err != nil {
			logger.Log.Error("append card entity", zap.Error(err))
			return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
		}
	}
	//-----------------------------------------------
	nodesLoginPasswordEntity, err := c.grpc.EntityGetList(ctx,
		&grpc.GetListEntityRequest{Type: constants.LoginPassword.ToString(),
			AccessToken: &grpc.Token{Token: token.AccessToken,
				UserId: token.UserID, CreatedAt: created, EndDateAt: endDate}})
	if err != nil {
		logger.Log.Error("entity get list", zap.Error(err))
		return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
	}
	for _, node := range nodesLoginPasswordEntity.Node {
		plaintext, err = utils.Decrypt(string(node.Data), secretKey)
		if err != nil {
			logger.Log.Error("decrypt", zap.Error(err))
			return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
		}
		var loginPassword model.LoginPassword
		err = json.Unmarshal([]byte(plaintext), &loginPassword)
		if err != nil {
			logger.Log.Error("unmarshal", zap.Error(err))
			return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
		}
		err = AppendLoginPasswordEntity(node, dataTblLoginPasswordPointer, loginPassword)
		if err != nil {
			logger.Log.Error("append login password entity", zap.Error(err))
			return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
		}
	}
	//-----------------------------------------------
	nodesBinary, err := c.grpc.FileGetList(ctx,
		&grpc.GetListBinaryRequest{AccessToken: &grpc.Token{Token: token.AccessToken,
			UserId: token.UserID, CreatedAt: created, EndDateAt: endDate}})
	if err != nil {
		logger.Log.Error("file get list", zap.Error(err))
		return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
	}

	for _, node := range nodesBinary.Node {
		err = AppendBinary(node, dataTblBinaryPointer)
		if err != nil {
			logger.Log.Error("append binary entity", zap.Error(err))
			return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
		}
	}
	//-----------------------------------------------
	return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, nil
}
