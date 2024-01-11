package ports

import (
	"context"
	"product-service/internal/core/entities"
)

// ProductService outlines the operations our product service must support
type ProductService interface {
	CreateProduct(ctx context.Context, product *entities.Product) (*entities.Product, error)
	GetProduct(ctx context.Context, id string) (*entities.Product, error)
	UpdateProduct(ctx context.Context, product *entities.Product) error
	DeleteProduct(ctx context.Context, id string) error
}
