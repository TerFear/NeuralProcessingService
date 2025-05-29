package app

import (
	grpcapp "gRPC_get_message/internal/app/grpc"
	"gRPC_get_message/internal/database/redis"
	"gRPC_get_message/internal/kafka"
	"gRPC_get_message/internal/services/prc"
	"go.uber.org/zap"
)

type App struct {
	GRPSrv *grpcapp.APP
}

func New(log *zap.Logger, grpcPort int, tokenTTL string) *App {

	r := redis.New(log)
	k := kafka.New(log)

	processer := prc.New(log, r, k)

	grpcApp := grpcapp.New(log, grpcPort, processer)

	return &App{
		GRPSrv: grpcApp,
	}
}
