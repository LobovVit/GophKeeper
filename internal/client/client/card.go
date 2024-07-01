package client

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/LobovVit/GophKeeper/internal/client/constants"
	"github.com/LobovVit/GophKeeper/internal/client/model"
	grpc "github.com/LobovVit/GophKeeper/internal/proto"
	"github.com/LobovVit/GophKeeper/pkg/logger"
	"github.com/LobovVit/GophKeeper/pkg/utils"
	"go.uber.org/zap"
)

// CardCreate - create card
func (c Client) CardCreate(ctx context.Context, name, description, password, paymentSystem, number, holder, cvc, endDate string, token model.Token) error {
	logger.Log.Info("card create ")

	intCvc, err := strconv.Atoi(cvc)
	if err != nil {
		logger.Log.Error("convert to int", zap.Error(err))
		return err
	}
	timeEndDate, err := time.Parse(constants.LayoutDate.ToString(), endDate)
	if err != nil {
		logger.Log.Error("parse", zap.Error(err))
		return err
	}
	card := model.Card{Name: name, Description: description, PaymentSystem: paymentSystem, Number: number, Holder: holder, EndDate: timeEndDate, CVC: intCvc}
	jsonCard, err := json.Marshal(card)
	if err != nil {
		logger.Log.Error("marshal", zap.Error(err))
		return err
	}

	secretKey := utils.AesKeySecureRandom([]byte(password))
	encryptCard, err := utils.Encrypt(string(jsonCard), secretKey)
	if err != nil {
		logger.Log.Error("encrypt", zap.Error(err))
		return err
	}
	createdToken := utils.ConvertTimeToTimestamp(token.CreatedAt)
	endDateToken := utils.ConvertTimeToTimestamp(token.EndDateAt)

	metadata := model.MetadataEntity{Name: name, Description: description, Type: constants.Card.ToString()}
	jsonMetadata, err := json.Marshal(metadata)
	if err != nil {
		return err
	}
	_, err = c.grpc.EntityCreate(ctx,
		&grpc.CreateEntityRequest{Data: []byte(encryptCard), Metadata: string(jsonMetadata),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		logger.Log.Error("entity create", zap.Error(err))
		return err
	}

	return nil
}

// CardDelete -  delete card
func (c Client) CardDelete(ctx context.Context, card []string, token model.Token) error {
	logger.Log.Info("card delete")

	createdToken := utils.ConvertTimeToTimestamp(token.CreatedAt)
	endDateToken := utils.ConvertTimeToTimestamp(token.EndDateAt)

	_, err := c.grpc.EntityDelete(ctx,
		&grpc.DeleteEntityRequest{Name: card[0], Type: constants.Card.ToString(),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID,
				CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		logger.Log.Error("entity delete", zap.Error(err))
		return err
	}

	return nil
}

// CardUpdate - update card
func (c Client) CardUpdate(ctx context.Context, name, passwordSecure, paymentSystem, number, holder, cvc, endDateCard string, token model.Token) error {
	logger.Log.Info("card update")

	intCvc, err := strconv.Atoi(cvc)
	if err != nil {
		logger.Log.Error("convert to int", zap.Error(err))
		return err
	}
	timeEndDate, err := time.Parse(constants.LayoutDate.ToString(), endDateCard)
	if err != nil {
		logger.Log.Error("parse", zap.Error(err))
		return err
	}
	card := model.Card{PaymentSystem: paymentSystem, Number: number, Holder: holder, CVC: intCvc, EndDate: timeEndDate}
	jsonCard, err := json.Marshal(card)
	if err != nil {
		logger.Log.Error("marshal", zap.Error(err))
		return err
	}

	secretKey := utils.AesKeySecureRandom([]byte(passwordSecure))
	encryptCard, err := utils.Encrypt(string(jsonCard), secretKey)
	if err != nil {
		logger.Log.Error("encrypt", zap.Error(err))
		return err
	}

	createdToken := utils.ConvertTimeToTimestamp(token.CreatedAt)
	endDateToken := utils.ConvertTimeToTimestamp(token.EndDateAt)

	_, err = c.grpc.EntityUpdate(ctx,
		&grpc.UpdateEntityRequest{Name: name, Data: []byte(encryptCard), Type: constants.Card.ToString(),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		logger.Log.Error("entity update", zap.Error(err))
		return err
	}

	return nil
}
