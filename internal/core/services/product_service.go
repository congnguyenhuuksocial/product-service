package services

import (
	"context"
	"go.uber.org/zap"
	"product-service/internal/core/entities"
	"product-service/internal/core/ports"
	"product-service/internal/infrastructure/eventstore"
	"product-service/internal/infrastructure/messagebus"
	"product-service/internal/infrastructure/messagebus/messages"
	"product-service/internal/infrastructure/repository"
	"time"
)

type productServiceImpl struct {
	productRepo ports.ProductRepository
	eventStore  ports.EventStore
	eventBus    ports.MessageBus
	logger      *zap.Logger
}

func NewProductService(
	repo *repository.ProductRepository,
	eventStore *eventstore.EventStore,
	eventBus *messagebus.KafkaBus,
	logger *zap.Logger,
) ports.ProductService {
	return &productServiceImpl{
		productRepo: repo,
		eventStore:  eventStore,
		eventBus:    eventBus,
		logger:      logger,
	}
}

func (s *productServiceImpl) CreateProduct(ctx context.Context, product *entities.Product) error {
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

	if err := s.eventBus.Publish(ctx, "product.created", messages.KafkaMessage{
		Key:   event.EventName(),
		Value: event.Marshal(),
	}); err != nil {
		s.logger.Error("failed to publish event", zap.Error(err))
	}
	return nil
}

func (s *productServiceImpl) GetProduct(ctx context.Context, id string) (*entities.Product, error) {
	return s.productRepo.FindByID(ctx, id)
}

func (s *productServiceImpl) UpdateProduct(ctx context.Context, product *entities.Product) error {
	return s.productRepo.Update(ctx, product)
}

func (s *productServiceImpl) DeleteProduct(ctx context.Context, id string) error {
	return s.productRepo.Delete(ctx, id)
}
