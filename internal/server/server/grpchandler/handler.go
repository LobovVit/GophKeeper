package grpchandler

import (
	"context"
	"database/sql"

	grpc "github.com/LobovVit/GophKeeper/internal/proto"
	"github.com/LobovVit/GophKeeper/internal/server/config"
	"github.com/LobovVit/GophKeeper/internal/server/domain/storage/repositories"
	"github.com/LobovVit/GophKeeper/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	database *sql.DB
	config   *config.Config
	user     *repositories.User
	file     *repositories.Files
	storage  *repositories.Storage
	entity   *repositories.Entity
	token    *repositories.Token
	grpc.UnimplementedGophKeeperServer
}

// NewHandler - creates a new grpc server instance
func NewHandler(db *sql.DB, config *config.Config, userRepository *repositories.User,
	binaryRepository *repositories.Files, storage *repositories.Storage, entityRepository *repositories.Entity, tokenRepository *repositories.Token) *Handler {
	return &Handler{database: db, config: config, user: userRepository, file: binaryRepository, storage: storage,
		entity: entityRepository, token: tokenRepository}
}

// Ping - checks the database connection
func (h *Handler) Ping(ctx context.Context, req *grpc.PingRequest) (*grpc.PingResponse, error) {
	logger.Log.Info("ping")
	var msg string
	err := h.database.Ping()
	if err != nil {
		msg = "unsuccessful database connection"
		logger.Log.Error("ping", zap.Error(err))
		return &grpc.PingResponse{Message: msg}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	msg = "successful database connection"
	logger.Log.Info(msg)
	return &grpc.PingResponse{Message: msg}, nil
}
