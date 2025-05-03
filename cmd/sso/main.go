package main

import (
	"gRPC_get_message/internal/app"
	"gRPC_get_message/internal/config"
	"gRPC_get_message/internal/logger/setuplogger"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.ReadConfig()

	log := setuplogger.InitLogger(cfg.Env)

	application := app.New(log, cfg.GRPC.Port, cfg.Token)

	go application.GRPSrv.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT) //выполнит сначала действия, а потом завершится

	sign := <-stop

	log.Info("Shutting down...", zap.String("signal", sign.String()))

	application.GRPSrv.Stop()
	log.Info("app graceful shutdown")

}
