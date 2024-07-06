package client

import (
	"context"

	"github.com/LobovVit/GophKeeper/internal/client/model"
	grpc "github.com/LobovVit/GophKeeper/internal/proto"
	"github.com/LobovVit/GophKeeper/pkg/logger"
	"github.com/LobovVit/GophKeeper/pkg/utils"
	"go.uber.org/zap"
)

// Authentication - auth user and return token
func (c Client) Authentication(ctx context.Context, username, password string) (model.Token, error) {
	logger.Log.Info("authentication")

	token := model.Token{}
	password, err := utils.HashPassword(password)
	if err != nil {
		logger.Log.Error("hash password", zap.Error(err))
		return token, err
	}
	authenticatedUser, err := c.grpc.Authentication(ctx, &grpc.AuthenticationRequest{Username: username, Password: password})
	if err != nil {
		logger.Log.Error("authentication", zap.Error(err))
		return token, err
	}

	createdToken, err := utils.ConvertTimestampToTime(authenticatedUser.AccessToken.CreatedAt)
	if err != nil {
		logger.Log.Error("convert timestamp to time", zap.Error(err))
		return token, err
	}
	endDateToken, err := utils.ConvertTimestampToTime(authenticatedUser.AccessToken.EndDateAt)
	if err != nil {
		logger.Log.Error("convert timestamp to time", zap.Error(err))
		return token, err
	}
	token = model.Token{AccessToken: authenticatedUser.AccessToken.Token, UserID: authenticatedUser.AccessToken.UserId,
		CreatedAt: createdToken, EndDateAt: endDateToken}

	return token, nil
}

// UserExist - check if user exist in db
func (c Client) UserExist(ctx context.Context, username string) (bool, error) {
	logger.Log.Info("user exist check")

	user, err := c.grpc.UserExist(ctx, &grpc.UserExistRequest{Username: username})
	if err != nil {
		logger.Log.Error("user exist", zap.Error(err))
		return user.Exist, err
	}
	return user.Exist, nil
}

// Registration - user registration
func (c Client) Registration(ctx context.Context, username, password string) (model.Token, error) {
	logger.Log.Info("registration")

	token := model.Token{}
	password, err := utils.HashPassword(password)
	if err != nil {
		logger.Log.Error("hash password", zap.Error(err))
		return token, err
	}
	registeredUser, err := c.grpc.Registration(ctx, &grpc.RegistrationRequest{Username: username, Password: password})
	if err != nil {
		logger.Log.Error("registration", zap.Error(err))
		return token, err
	}
	createdToken, err := utils.ConvertTimestampToTime(registeredUser.AccessToken.CreatedAt)
	if err != nil {
		logger.Log.Error("convert timestamp to time", zap.Error(err))
		return token, err
	}
	endDateToken, err := utils.ConvertTimestampToTime(registeredUser.AccessToken.EndDateAt)
	if err != nil {
		logger.Log.Error("convert timestamp to time", zap.Error(err))
		return token, err
	}
	token = model.Token{AccessToken: registeredUser.AccessToken.Token, UserID: registeredUser.AccessToken.UserId,
		CreatedAt: createdToken, EndDateAt: endDateToken}

	return token, nil
}
