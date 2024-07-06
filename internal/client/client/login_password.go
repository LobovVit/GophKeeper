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

// LoginPasswordCreate - add login-password
func (c Client) LoginPasswordCreate(ctx context.Context, name, description, passwordSecure, login, password string, token model.Token) error {
	logger.Log.Info("login password create")

	loginPassword := model.LoginPassword{Login: login, Password: password}
	jsonLoginPassword, err := json.Marshal(loginPassword)
	if err != nil {
		logger.Log.Error("marshal", zap.Error(err))
		return err
	}

	secretKey := utils.AesKeySecureRandom([]byte(passwordSecure))
	encryptLoginPassword, err := utils.Encrypt(string(jsonLoginPassword), secretKey)
	if err != nil {
		logger.Log.Error("encrypt", zap.Error(err))
		return err
	}
	createdToken := utils.ConvertTimeToTimestamp(token.CreatedAt)

	endDateToken := utils.ConvertTimeToTimestamp(token.EndDateAt)

	metadata := model.MetadataEntity{Name: name, Description: description, Type: constants.LoginPassword.ToString()}
	jsonMetadata, err := json.Marshal(metadata)
	if err != nil {
		return err
	}
	_, err = c.grpc.EntityCreate(ctx,
		&grpc.CreateEntityRequest{Data: []byte(encryptLoginPassword), Metadata: string(jsonMetadata),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		logger.Log.Error("entity create", zap.Error(err))
		return err
	}

	return nil
}

// LoginPasswordDelete - delete login-password
func (c Client) LoginPasswordDelete(ctx context.Context, loginPassword []string, token model.Token) error {
	logger.Log.Info("login password delete")

	createdToken := utils.ConvertTimeToTimestamp(token.CreatedAt)
	endDateToken := utils.ConvertTimeToTimestamp(token.EndDateAt)

	_, err := c.grpc.EntityDelete(ctx,
		&grpc.DeleteEntityRequest{Name: loginPassword[0], Type: constants.LoginPassword.ToString(),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID,
				CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		logger.Log.Error("entity delete", zap.Error(err))
		return err
	}

	return nil
}

// LoginPasswordUpdate - update login-password
func (c Client) LoginPasswordUpdate(ctx context.Context, name, passwordSecure, login, password string, token model.Token) error {
	logger.Log.Info("login password update")

	loginPassword := model.LoginPassword{Login: login, Password: password}
	jsonLoginPassword, err := json.Marshal(loginPassword)
	if err != nil {
		logger.Log.Error("marshal", zap.Error(err))
		return err
	}

	secretKey := utils.AesKeySecureRandom([]byte(passwordSecure))
	encryptLoginPassword, err := utils.Encrypt(string(jsonLoginPassword), secretKey)
	if err != nil {
		logger.Log.Error("encrypt", zap.Error(err))
		return err
	}
	createdToken := utils.ConvertTimeToTimestamp(token.CreatedAt)

	endDateToken := utils.ConvertTimeToTimestamp(token.EndDateAt)

	_, err = c.grpc.EntityUpdate(ctx,
		&grpc.UpdateEntityRequest{Name: name, Data: []byte(encryptLoginPassword), Type: constants.LoginPassword.ToString(),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		logger.Log.Error("entity update", zap.Error(err))
		return err
	}

	return nil
}
