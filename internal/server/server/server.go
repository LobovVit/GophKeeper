package server

import (
	"context"
	"database/sql"
	"net"
	"sync"

	pb "github.com/LobovVit/GophKeeper/internal/proto"
	"github.com/LobovVit/GophKeeper/internal/server/config"
	"github.com/LobovVit/GophKeeper/internal/server/domain/storage/repositories"
	"github.com/LobovVit/GophKeeper/internal/server/server/grpchandler"
	"github.com/LobovVit/GophKeeper/pkg/logger"
	"github.com/LobovVit/GophKeeper/pkg/postgresql"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

// Server - structure containing a server instance
type Server struct {
	config  *config.Config
	storage *sql.DB
	wg      sync.WaitGroup
}

// New - method to create server instance
func New(ctx context.Context, config *config.Config) (*Server, error) {
	stor, err := postgresql.NewConn(ctx, config.DSN)
	if err != nil {
		return nil, err
	}
	return &Server{config: config, storage: stor}, nil
}

// Run - method to start server instance
func (a *Server) Run(ctx context.Context) error {
	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", a.config.HostGRPC)
	if err != nil {
		logger.Log.Error("listen", zap.Error(err))
		return err
	}
	userRepository := repositories.NewUserRepo(a.storage)
	binaryRepository := repositories.NewFileRepo(a.storage)
	fileStorage := repositories.New(a.config.Files)
	entityRepository := repositories.NewEntityRepo(a.storage)
	tokenRepository := repositories.NewTokenRepo(a.storage)
	handlerGrpc := grpchandler.NewHandler(a.storage, a.config, userRepository, binaryRepository, &fileStorage, entityRepository, tokenRepository)
	pb.RegisterGophKeeperServer(grpcServer, handlerGrpc)

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		logger.Log.Info("Starting grpc server", zap.String("address", a.config.HostGRPC))
		return grpcServer.Serve(lis)
	})

	a.wg.Add(1)
	go (func() {
		<-gCtx.Done()
		a.Shutdown(grpcServer)
	})()
	if err = g.Wait(); err != nil { //
		logger.Log.Error("server:", zap.Error(err))
	}
	a.wg.Wait()
	return nil
}

// Shutdown - method that implements saving the server state when shutting down
func (a *Server) Shutdown(srvGRPC *grpc.Server) {
	defer a.wg.Done()
	srvGRPC.GracefulStop()
	logger.Log.Info("grpc server shutdown ok")
}
