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

// FileDownload - check token, save record, upload file on client
func (h *Handler) FileDownload(ctx context.Context, req *grpc.DownloadBinaryRequest) (*grpc.DownloadBinaryResponse, error) {
	logger.Log.Info("file download")

	endDateToken, err := h.token.GetEndDateToken(ctx, req.AccessToken.Token)
	if err != nil {
		logger.Log.Error("get end date token", zap.Error(err))
		return &grpc.DownloadBinaryResponse{}, err
	}
	valid := h.token.Validate(endDateToken)
	if !valid {
		logger.Log.Error("validate", zap.Error(errors.ErrNotValidateToken))
		return &grpc.DownloadBinaryResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	FileData := &model.FileRequest{}
	FileData.UserID = req.AccessToken.UserId
	FileData.Name = req.Name

	exists, err := h.file.FileExists(ctx, FileData)
	if err != nil {
		logger.Log.Error("file exists", zap.Error(err))
		return &grpc.DownloadBinaryResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	if !exists {
		err = errors.ErrFileNotExists
		logger.Log.Error("file not exists", zap.Error(err))
		return &grpc.DownloadBinaryResponse{}, status.Errorf(
			codes.AlreadyExists, err.Error(),
		)
	}

	data, err := utils.DownloadFile(h.config.Files, req.AccessToken.UserId, req.Name)
	if err != nil {
		logger.Log.Error("download file", zap.Error(err))
		return &grpc.DownloadBinaryResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	return &grpc.DownloadBinaryResponse{Data: data}, nil
}

// FileGetList - checks tokens and get list records file
func (h *Handler) FileGetList(ctx context.Context, req *grpc.GetListBinaryRequest) (*grpc.GetListBinaryResponse, error) {
	logger.Log.Info("file get list")

	endDateToken, err := h.token.GetEndDateToken(ctx, req.AccessToken.Token)
	if err != nil {
		logger.Log.Error("get end date token", zap.Error(err))
		return &grpc.GetListBinaryResponse{}, err
	}
	valid := h.token.Validate(endDateToken)
	if !valid {
		logger.Log.Error("validate", zap.Error(errors.ErrNotValidateToken))
		return &grpc.GetListBinaryResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	ListFile, err := h.file.GetListFile(ctx, req.AccessToken.UserId)
	if err != nil {
		logger.Log.Error("get list file", zap.Error(err))
		return &grpc.GetListBinaryResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	list := model.GetListFile(ListFile)

	return &grpc.GetListBinaryResponse{Node: list}, nil
}

// FileRemove - checks the validity of the token, delete record, remove file on server
func (h *Handler) FileRemove(ctx context.Context, req *grpc.DeleteBinaryRequest) (*grpc.DeleteBinaryResponse, error) {
	logger.Log.Info("file remove")

	endDateToken, err := h.token.GetEndDateToken(ctx, req.AccessToken.Token)
	if err != nil {
		logger.Log.Error("get end date token", zap.Error(err))
		return &grpc.DeleteBinaryResponse{}, err
	}
	valid := h.token.Validate(endDateToken)
	if !valid {
		logger.Log.Error("validate", zap.Error(errors.ErrNotValidateToken))
		return &grpc.DeleteBinaryResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	FileData := &model.FileRequest{}
	FileData.UserID = req.AccessToken.UserId
	FileData.Name = req.Name

	BinaryId, err := h.file.DeleteFile(ctx, FileData)
	if err != nil {
		logger.Log.Error("delete file", zap.Error(err))
		return &grpc.DeleteBinaryResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	err = utils.RemoveFile(h.config.Files, req.AccessToken.UserId, req.Name)
	if err != nil {
		logger.Log.Error("remove file", zap.Error(err))
		return &grpc.DeleteBinaryResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	return &grpc.DeleteBinaryResponse{Id: BinaryId}, nil
}

// FileUpload - checks the validity of the token, upload file on client
func (h *Handler) FileUpload(ctx context.Context, req *grpc.UploadBinaryRequest) (*grpc.UploadBinaryResponse, error) {
	logger.Log.Info("file upload")

	endDateToken, err := h.token.GetEndDateToken(ctx, req.AccessToken.Token)
	if err != nil {
		logger.Log.Error("get end date token", zap.Error(err))
		return &grpc.UploadBinaryResponse{}, err
	}
	valid := h.token.Validate(endDateToken)
	if !valid {
		logger.Log.Error("validate", zap.Error(errors.ErrNotValidateToken))
		return &grpc.UploadBinaryResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	FileData := &model.FileRequest{}
	FileData.UserID = req.AccessToken.UserId
	FileData.Name = req.Name

	exists, err := h.file.FileExists(ctx, FileData)
	if err != nil {
		logger.Log.Error("file exists", zap.Error(err))
		return &grpc.UploadBinaryResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	if exists {
		err = errors.ErrNameAlreadyExists
		logger.Log.Error("name already exists", zap.Error(err))
		return &grpc.UploadBinaryResponse{}, status.Errorf(
			codes.AlreadyExists, err.Error(),
		)
	}

	UploadFile, err := h.file.UploadFile(ctx, FileData)
	if err != nil {
		logger.Log.Error("upload file", zap.Error(err))
		return &grpc.UploadBinaryResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	err = utils.UploadFile(h.config.Files, req.AccessToken.UserId, req.Name, req.Data)
	if err != nil {
		logger.Log.Error("upload file", zap.Error(err))
		return &grpc.UploadBinaryResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	return &grpc.UploadBinaryResponse{Name: UploadFile.Name}, nil
}
