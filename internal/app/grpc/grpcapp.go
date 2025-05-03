package grpc

import (
	"fmt"
	"gRPC_get_message/internal/grpc/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

// структура приложения
type APP struct {
	GRPCsrv *grpc.Server
	log     *zap.Logger
	port    int
}

// регистрация нового gRPC
func New(log *zap.Logger, port int) *APP {
	gRPC := grpc.NewServer()

	server.Register(gRPC)
	log.Info("gRPC server created")
	return &APP{
		log:     log,
		port:    port,
		GRPCsrv: gRPC,
	}
}

func (app *APP) MustRun() {
	err := app.Start()

	if err != nil {
		app.log.Error("Error starting gRPC server", zap.Error(err))
	}

}

func (app *APP) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", app.port))

	if err != nil {
		app.log.Error("Failed to listen", zap.Int("port", app.port), zap.Error(err))
		return err
	}
	app.log.Info("gRPC server listening", zap.Int("port", app.port))

	if err := app.GRPCsrv.Serve(lis); err != nil {
		app.log.Error("Failed to serve", zap.Error(err))
		return err
	}
	return nil
}

func (app *APP) Stop() {
	app.log.Info("gRPC server stopped")
	app.GRPCsrv.GracefulStop()
	app.log.Info("gRPC server stopped")
}
