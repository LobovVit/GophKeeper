package grpchandler

import (
	"context"

	grpc "github.com/LobovVit/GophKeeper/internal/proto"
	"github.com/LobovVit/GophKeeper/internal/server/domain/model"
	"github.com/LobovVit/GophKeeper/internal/server/domain/storage/errors"
	"github.com/LobovVit/GophKeeper/pkg/logger"
	"github.com/LobovVit/GophKeeper/pkg/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Authentication - user authentication, create access token
func (h *Handler) Authentication(ctx context.Context, req *grpc.AuthenticationRequest) (*grpc.AuthenticationResponse, error) {
	logger.Log.Info("authentication")
	UserData := &model.UserRequest{
		Username: req.Username,
		Password: req.Password,
	}

	authenticatedUser, err := h.user.Authentication(ctx, UserData)
	if err != nil {
		logger.Log.Error("authentication", zap.Error(err))
		return &grpc.AuthenticationResponse{}, status.Errorf(
			codes.Unauthenticated, err.Error(),
		)
	}
	user := model.GetUserData(authenticatedUser)

	token, err := h.token.Create(ctx, user.UserId, h.config.TokenLifetime)
	if err != nil {
		logger.Log.Error("token create", zap.Error(err))
		return &grpc.AuthenticationResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	createdToken := utils.ConvertTimeToTimestamp(token.CreatedAt)
	endDateToken := utils.ConvertTimeToTimestamp(token.EndDateAt)

	return &grpc.AuthenticationResponse{AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}}, nil
}

// UserExist - check user exist
func (h *Handler) UserExist(ctx context.Context, req *grpc.UserExistRequest) (*grpc.UserExistResponse, error) {
	logger.Log.Info("user exist check")
	exist, err := h.user.UserExists(ctx, req.Username)
	if err != nil {
		logger.Log.Error("user exists", zap.Error(err))
		return &grpc.UserExistResponse{Exist: false}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	return &grpc.UserExistResponse{Exist: exist}, nil
}

// Registration - registration new user, create access token
func (h *Handler) Registration(ctx context.Context, req *grpc.RegistrationRequest) (*grpc.RegistrationResponse, error) {
	logger.Log.Info("registration")

	UserData := &model.UserRequest{}
	UserData.Username = req.Username
	UserData.Password = req.Password

	exists, err := h.user.UserExists(ctx, UserData.Username)
	if err != nil {
		logger.Log.Error("user exists", zap.Error(err))
		return &grpc.RegistrationResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	if exists {
		err = errors.ErrUsernameAlreadyExists
		logger.Log.Error("user exists", zap.Error(err))
		return &grpc.RegistrationResponse{}, status.Errorf(
			codes.AlreadyExists, err.Error(),
		)
	}
	registeredUser, err := h.user.Registration(ctx, UserData)
	if err != nil {
		logger.Log.Error("user registration", zap.Error(err))
		return &grpc.RegistrationResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	user := model.GetUserData(registeredUser)

	token, err := h.token.Create(ctx, user.UserId, h.config.TokenLifetime)
	if err != nil {
		logger.Log.Error("token create", zap.Error(err))
		return &grpc.RegistrationResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	createdToken := utils.ConvertTimeToTimestamp(token.CreatedAt)
	endDateToken := utils.ConvertTimeToTimestamp(token.EndDateAt)

	err = utils.CreateStorageUser(h.config.Files, token.UserID)
	if err != nil {
		logger.Log.Error("create constants user", zap.Error(err))
		return &grpc.RegistrationResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	return &grpc.RegistrationResponse{AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID,
		CreatedAt: createdToken, EndDateAt: endDateToken}}, nil
}
