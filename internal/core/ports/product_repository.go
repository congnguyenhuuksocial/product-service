// File: /internal/core/ports/productrepository.go

package ports

import (
	"context"
	"product-service/internal/core/entities"
)

// ProductRepository defines the expected behaviour from a product repository
type ProductRepository interface {
	Create(ctx context.Context, product *entities.Product) error
	FindByID(ctx context.Context, id string) (*entities.Product, error)
	Update(ctx context.Context, product *entities.Product) error
	Delete(ctx context.Context, id string) error
}
