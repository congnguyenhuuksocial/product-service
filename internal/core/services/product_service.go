// File: /internal/core/services/productservice.go

package services

import (
	"context"
	"product-service/internal/core/entities"
	"product-service/internal/core/ports"
	"time"
)

type productServiceImpl struct {
	productRepo ports.ProductRepository
	eventStore  ports.EventStore
}

func NewProductService(repo ports.ProductRepository, eventStore ports.EventStore) ports.ProductService {
	return &productServiceImpl{
		productRepo: repo,
		eventStore:  eventStore,
	}
}

func (s *productServiceImpl) CreateProduct(ctx context.Context, product *entities.Product) error {
	// Implement your business logic here
	// For instance, validate the product, calculate any fields or apply business rules
	if err := s.productRepo.Create(ctx, product); err != nil {
		return err
	}

	event := entities.ProductCreatedEvent{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Created:     time.Now().UTC(),
	}

	if err := s.eventStore.AppendEvent(ctx, event); err != nil {
		return err
	}
	return nil
}

func (s *productServiceImpl) GetProduct(ctx context.Context, id string) (*entities.Product, error) {
	// Business logic can also go here, e.g. caching strategy
	return s.productRepo.FindByID(ctx, id)
}

func (s *productServiceImpl) UpdateProduct(ctx context.Context, product *entities.Product) error {
	// Include logic for update such as validation or rules
	return s.productRepo.Update(ctx, product)
}

func (s *productServiceImpl) DeleteProduct(ctx context.Context, id string) error {
	// Possible checks before deletion
	return s.productRepo.Delete(ctx, id)
}
