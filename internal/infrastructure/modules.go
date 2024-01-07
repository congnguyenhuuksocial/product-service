package infrastructure

import (
	"go.uber.org/fx"
	"product-service/internal/infrastructure/eventstore"
	"product-service/internal/infrastructure/messagebus"
	"product-service/internal/infrastructure/repository"
)

var Module = fx.Options(
	fx.Provide(eventstore.NewEventStore),
	fx.Provide(messagebus.NewKafkaBus),
	fx.Provide(repository.NewProductRepository),
)
