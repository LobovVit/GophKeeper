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

// TextCreate - add text
func (c Client) TextCreate(ctx context.Context, name, description, password, plaintext string, token model.Token) error {
	logger.Log.Info("text create")

	secretKey := utils.AesKeySecureRandom([]byte(password))
	encryptText, err := utils.Encrypt(plaintext, secretKey)
	if err != nil {
		logger.Log.Error("encrypt", zap.Error(err))
		return err
	}
	createdToken := utils.ConvertTimeToTimestamp(token.CreatedAt)
	endDateToken := utils.ConvertTimeToTimestamp(token.EndDateAt)

	metadata := model.MetadataEntity{Name: name, Description: description, Type: constants.Text.ToString()}
	jsonMetadata, err := json.Marshal(metadata)
	if err != nil {
		return err
	}
	_, err = c.grpc.EntityCreate(ctx,
		&grpc.CreateEntityRequest{Data: []byte(encryptText), Metadata: string(jsonMetadata),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		logger.Log.Error("entity create", zap.Error(err))
		return err
	}

	return nil
}

// TextDelete - delete text
func (c Client) TextDelete(ctx context.Context, text []string, token model.Token) error {
	logger.Log.Info("text delete")

	createdToken := utils.ConvertTimeToTimestamp(token.CreatedAt)

	endDateToken := utils.ConvertTimeToTimestamp(token.EndDateAt)

	_, err := c.grpc.EntityDelete(ctx,
		&grpc.DeleteEntityRequest{Name: text[0], Type: constants.Text.ToString(),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID,
				CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		logger.Log.Error("entity delete", zap.Error(err))
		return err
	}

	return nil
}

// TextUpdate - update text
func (c Client) TextUpdate(ctx context.Context, name, passwordSecure, text string, token model.Token) error {
	logger.Log.Info("text update")

	secretKey := utils.AesKeySecureRandom([]byte(passwordSecure))
	encryptText, err := utils.Encrypt(text, secretKey)
	if err != nil {
		logger.Log.Error("encrypt", zap.Error(err))
		return err
	}

	createdToken := utils.ConvertTimeToTimestamp(token.CreatedAt)
	endDateToken := utils.ConvertTimeToTimestamp(token.EndDateAt)

	_, err = c.grpc.EntityUpdate(ctx,
		&grpc.UpdateEntityRequest{Name: name, Data: []byte(encryptText), Type: constants.Text.ToString(),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		logger.Log.Error("entity update", zap.Error(err))
		return err
	}

	return nil
}
