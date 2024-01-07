package core

import (
	"go.uber.org/fx"
	"product-service/internal/core/services"
)

var Module = fx.Options(
	fx.Provide(services.NewProductService),
)
