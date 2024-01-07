package adapter

import (
	"go.uber.org/fx"
	"product-service/internal/adapter/database"
	"product-service/internal/adapter/elasticsearch"
	healthv1 "product-service/internal/adapter/grpc/healthcheck/v1"
	productv1 "product-service/internal/adapter/grpc/product/v1"
	"product-service/internal/adapter/kafka"
	"product-service/internal/adapter/redis"
)

var Module = fx.Options(
	fx.Provide(
		productv1.NewProductService,
		healthv1.NewHealthCheckService,
	),
	fx.Provide(database.NewDatabase),
	fx.Provide(redis.NewClient),
	fx.Provide(elasticsearch.NewSearchClient),
	fx.Provide(kafka.NewKafka),
)
