package application

import (
	"go.uber.org/fx"
	"product-service/internal/application/commandhandlers"
	"product-service/internal/application/queryhandlers"
)

var Module = fx.Options(
	fx.Provide(commandhandlers.NewCreateProductCommandHandler),
	fx.Provide(
		queryhandlers.NewGetProductQueryHandler,
	),
)
