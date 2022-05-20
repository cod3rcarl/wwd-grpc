package main

import (
	"fmt"
	"log"

	"go.uber.org/zap"

	baseApp "github.com/cod3rcarl/wwd-grpc/base/app"
	baseLogger "github.com/cod3rcarl/wwd-grpc/base/logger"
	basePgx "github.com/cod3rcarl/wwd-grpc/base/pgx"

	"github.com/cod3rcarl/wwd-grpc/pkg/server"
	"github.com/cod3rcarl/wwd-grpc/pkg/storage"
	wwd "github.com/cod3rcarl/wwd-grpc/pkg/wwdatabase"
)

func main() {
	logger := createLogger()

	app := baseApp.NewApp(
		baseApp.WithLogger(logger),
	)

	// main business logic
	store := createStorage(logger)
	wwdService := wwd.NewService(
		wwd.WithLogger(logger),
		wwd.WithStorage(store),
	)

	grpcSrv := CreateGrpcServer(logger, wwdService)

	go func() {
		if err := grpcSrv.ListenAndServe(); err != nil {
			logger.Error("error running grpc server server", zap.Error(err))

			// a fatal log would call os.Exit(1), so cancel the context to trigger graceful shutdown
			app.Cancel()
		}
	}()

	serviceShutdownOrder := []baseApp.Stopper{
		grpcSrv,
	}

	app.HandleGracefulShutdown(serviceShutdownOrder)
}

func createLogger() *zap.Logger {
	cfg, err := baseLogger.ReadConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to read logger config, error=%v", err))
	}

	logger, err := baseLogger.NewLogger(cfg)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to create logger, error=%v", err))
	}

	return logger
}

func createStorage(logger *zap.Logger) *storage.Service {
	cfg, err := basePgx.ReadConfig()
	if err != nil {
		logger.Fatal(
			"failed to read config",
			zap.String("service", "storage"),
			zap.Error(err),
		)
	}

	base, err := basePgx.NewPgx(cfg, basePgx.WithLogger(logger))
	if err != nil {
		logger.Fatal(
			"failed to create service",
			zap.String("service", "storage"),
			zap.Error(err),
		)
	}

	return storage.NewStorage(storage.WithBase(base))
}

func CreateGrpcServer(
	logger *zap.Logger,
	wwdService wwd.ServiceInterface,
) *server.Service {
	cfg, err := server.ReadConfig()
	if err != nil {
		logger.Fatal(
			"failed to read config",
			zap.String("service", "grpc-server"),
			zap.Error(err),
		)
	}

	return server.NewServer(
		cfg,
		logger,
		server.WithClient(wwdService),
	)
}
