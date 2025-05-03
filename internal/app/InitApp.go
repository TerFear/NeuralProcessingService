package app

import (
	grpcapp "gRPC_get_message/internal/app/grpc"
	"go.uber.org/zap"
)

type App struct {
	GRPSrv *grpcapp.APP
}

func New(log *zap.Logger, grpcPort int, tokenTTL string) *App {
	//TODO инициализация хранилище

	//TODO инициализация server сервер

	//TODO инициализация gRCP сервер

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPSrv: grpcApp,
	}
}
