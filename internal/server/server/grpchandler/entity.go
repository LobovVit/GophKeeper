package grpchandler

import (
	"context"
	"encoding/json"

	grpc "github.com/LobovVit/GophKeeper/internal/proto"
	"github.com/LobovVit/GophKeeper/internal/server/domain/model"
	"github.com/LobovVit/GophKeeper/internal/server/domain/storage/errors"
	"github.com/LobovVit/GophKeeper/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// EntityGetList - checks the validity of token and get list records (text, bank card or login-password)
func (h *Handler) EntityGetList(ctx context.Context, req *grpc.GetListEntityRequest) (*grpc.GetListEntityResponse, error) {
	logger.Log.Info("Get list entity")

	endDateToken, err := h.token.GetEndDateToken(ctx, req.AccessToken.Token)
	if err != nil {
		logger.Log.Error("get end date token", zap.Error(err))
		return &grpc.GetListEntityResponse{}, err
	}
	valid := h.token.Validate(endDateToken)
	if !valid {
		logger.Log.Error("validate", zap.Error(errors.ErrNotValidateToken))
		return &grpc.GetListEntityResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	ListEntity, err := h.entity.GetList(ctx, req.AccessToken.UserId, req.Type)
	if err != nil {
		logger.Log.Error("get list", zap.Error(err))
		return &grpc.GetListEntityResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	list, err := model.GetListEntity(ListEntity)
	if err != nil {
		logger.Log.Error("get list entity", zap.Error(err))
		return &grpc.GetListEntityResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	return &grpc.GetListEntityResponse{Node: list}, nil
}

// EntityCreate - check token and save record
func (h *Handler) EntityCreate(ctx context.Context, req *grpc.CreateEntityRequest) (*grpc.CreateEntityResponse, error) {
	logger.Log.Info("entity create")

	endDateToken, err := h.token.GetEndDateToken(ctx, req.AccessToken.Token)
	if err != nil {
		logger.Log.Error("end date token", zap.Error(err))
		return &grpc.CreateEntityResponse{}, err
	}
	valid := h.token.Validate(endDateToken)
	if !valid {
		logger.Log.Error("validate token", zap.Error(errors.ErrNotValidateToken))
		return &grpc.CreateEntityResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	var metadata model.MetadataEntity
	err = json.Unmarshal([]byte(req.Metadata), &metadata)
	if err != nil {
		logger.Log.Error("Unmarshal metadata", zap.Error(err))
		return &grpc.CreateEntityResponse{}, err
	}

	EntityData := &model.CreateEntityRequest{}
	EntityData.UserID = req.AccessToken.UserId
	EntityData.Data = req.Data
	EntityData.Metadata = metadata
	if metadata.Name == "" {
		err := errors.ErrNoMetadataSet
		logger.Log.Error("no metadata set", zap.Error(err))
		return &grpc.CreateEntityResponse{}, status.Errorf(
			codes.InvalidArgument, err.Error(),
		)
	}

	exists, err := h.entity.Exists(ctx, EntityData)
	if err != nil {
		logger.Log.Error("Check exists", zap.Error(err))
		return &grpc.CreateEntityResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	if exists {
		err = errors.ErrNameAlreadyExists
		logger.Log.Error("already exists", zap.Error(err))
		return &grpc.CreateEntityResponse{}, status.Errorf(
			codes.AlreadyExists, err.Error(),
		)
	}

	CreatedEntityID, err := h.entity.Create(ctx, EntityData)
	if err != nil {
		logger.Log.Error("create entity", zap.Error(err))
		return &grpc.CreateEntityResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	return &grpc.CreateEntityResponse{Id: CreatedEntityID}, nil
}

// EntityDelete - check token and delete record
func (h *Handler) EntityDelete(ctx context.Context, req *grpc.DeleteEntityRequest) (*grpc.DeleteEntityResponse, error) {
	logger.Log.Info("delete entity")

	endDateToken, err := h.token.GetEndDateToken(ctx, req.AccessToken.Token)
	if err != nil {
		logger.Log.Error("end date token", zap.Error(err))
		return &grpc.DeleteEntityResponse{}, err
	}
	valid := h.token.Validate(endDateToken)
	if !valid {
		logger.Log.Error("validate token", zap.Error(errors.ErrNotValidateToken))
		return &grpc.DeleteEntityResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	DeletedEntityID, err := h.entity.Delete(ctx, req.AccessToken.UserId, req.Name, req.Type)
	if err != nil {
		logger.Log.Error("delete entity", zap.Error(err))
		return &grpc.DeleteEntityResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	return &grpc.DeleteEntityResponse{Id: DeletedEntityID}, nil
}

// EntityUpdate - checks token and update record
func (h *Handler) EntityUpdate(ctx context.Context, req *grpc.UpdateEntityRequest) (*grpc.UpdateEntityResponse, error) {
	logger.Log.Info("entity update")

	endDateToken, err := h.token.GetEndDateToken(ctx, req.AccessToken.Token)
	if err != nil {
		logger.Log.Error("end date token", zap.Error(err))
		return &grpc.UpdateEntityResponse{}, err
	}
	valid := h.token.Validate(endDateToken)
	if !valid {
		logger.Log.Error("validate token", zap.Error(errors.ErrNotValidateToken))
		return &grpc.UpdateEntityResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	UpdatedEntityID, err := h.entity.Update(ctx, req.AccessToken.UserId, req.Name, req.Type, req.Data)
	if err != nil {
		logger.Log.Error("update entity", zap.Error(err))
		return &grpc.UpdateEntityResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	return &grpc.UpdateEntityResponse{Id: UpdatedEntityID}, nil
}
