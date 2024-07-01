package client

import (
	"context"
	"fmt"

	"github.com/LobovVit/GophKeeper/internal/client/config"
	pb "github.com/LobovVit/GophKeeper/internal/proto"
	"github.com/LobovVit/GophKeeper/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	Cfg  *config.Config
	grpc pb.GophKeeperClient
	pb.UnimplementedGophKeeperServer
}

func New(config *config.Config) (*Client, error) {
	client := Client{Cfg: config}
	conn, err := grpc.NewClient(
		config.HostGRPC,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Log.Error("grpc new client", zap.Error(err))
		return nil, fmt.Errorf("grpc new client: %w", err)
	}
	client.grpc = pb.NewGophKeeperClient(conn)
	return &client, nil
}

// Ping - ping
func (c Client) Ping(ctx context.Context) (string, error) {
	logger.Log.Info("ping")

	msg, err := c.grpc.Ping(ctx, &pb.PingRequest{})
	if err != nil {
		logger.Log.Error("ping", zap.Error(err))
		return "", err
	}

	return msg.Message, nil
}
