package container

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"product-service/internal/adapter"
	healthv1 "product-service/internal/adapter/grpc/healthcheck/v1"
	productv1 "product-service/internal/adapter/grpc/product/v1"
	"product-service/internal/application"
	"product-service/internal/core"
	"product-service/internal/infrastructure"
	"product-service/pkg/config"
	"product-service/pkg/logger"
)

var Modules = fx.Options(
	infrastructure.Module,
	core.Module,
	application.Module,
	adapter.Module,
	config.Module,
	logger.Module,
	fx.Invoke(
		RegisterGrpcHook,
	),
)

func RegisterGrpcHook(
	lc fx.Lifecycle,
	conf *config.Config,
	log *zap.Logger,
	health *healthv1.HealthCheckService,
	product *productv1.ProductGrpcService,
) {
	grpcConf := conf.GRPC
	grpcServer := grpc.NewServer()
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Start grpc server")
			address := fmt.Sprintf(":%d", grpcConf.Port)
			lis, err := net.Listen("tcp", address)
			if err != nil {
				return err
			}
			reflection.Register(grpcServer)
			healthv1.RegisterHealthServer(grpcServer, health)
			productv1.RegisterProductServiceServer(grpcServer, product)

			fatalErr := make(chan error)
			go func() {
				if err := grpcServer.Serve(lis); err != nil {
					fatalErr <- err
				}
			}()

			go func() {
				defer close(fatalErr)
				_err := <-fatalErr
				log.Error("Error when start grpc server", zap.Error(_err))
			}()
			return nil
		},
		OnStop: func(_ context.Context) error {
			log.Info("Stop grpc server")
			grpcServer.GracefulStop()
			return nil
		},
	})
}
