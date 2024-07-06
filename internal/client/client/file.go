package client

import (
	"context"

	"github.com/LobovVit/GophKeeper/internal/client/model"
	grpc "github.com/LobovVit/GophKeeper/internal/proto"
	"github.com/LobovVit/GophKeeper/pkg/logger"
	"github.com/LobovVit/GophKeeper/pkg/utils"
	"go.uber.org/zap"
)

// FileUpload - upload file
func (c Client) FileUpload(ctx context.Context, name string, password string, file []byte, token model.Token) (string, error) {
	logger.Log.Info("file upload")

	secretKey := utils.AesKeySecureRandom([]byte(password))
	encryptFile, err := utils.Encrypt(string(file), secretKey)
	if err != nil {
		logger.Log.Error("encrypt", zap.Error(err))
		return "", err
	}
	createdToken := utils.ConvertTimeToTimestamp(token.CreatedAt)

	endDateToken := utils.ConvertTimeToTimestamp(token.EndDateAt)

	uploadFile, err := c.grpc.FileUpload(ctx,
		&grpc.UploadBinaryRequest{Name: name, Data: []byte(encryptFile),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID,
				CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		logger.Log.Error("file upload", zap.Error(err))
		return "", err
	}

	return uploadFile.Name, nil
}

// FileDownload - download file
func (c Client) FileDownload(ctx context.Context, name string, password string, token model.Token) ([]byte, error) {
	logger.Log.Info("file download")

	secretKey := utils.AesKeySecureRandom([]byte(password))
	createdToken := utils.ConvertTimeToTimestamp(token.CreatedAt)
	endDateToken := utils.ConvertTimeToTimestamp(token.EndDateAt)

	downloadFile, err := c.grpc.FileDownload(ctx,
		&grpc.DownloadBinaryRequest{Name: name, AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID,
			CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		logger.Log.Error("file download", zap.Error(err))
		return nil, err
	}

	file, err := utils.Decrypt(string(downloadFile.Data), secretKey)
	if err != nil {
		logger.Log.Error("decrypt", zap.Error(err))
		return nil, err
	}

	return []byte(file), nil
}

// FileRemove - delete file
func (c Client) FileRemove(ctx context.Context, binary []string, token model.Token) error {
	logger.Log.Info("file remove")

	createdToken := utils.ConvertTimeToTimestamp(token.CreatedAt)
	endDateToken := utils.ConvertTimeToTimestamp(token.EndDateAt)
	_, err := c.grpc.FileRemove(ctx,
		&grpc.DeleteBinaryRequest{Name: binary[0], AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID,
			CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		logger.Log.Error("file remove", zap.Error(err))
		return err
	}

	return nil
}
