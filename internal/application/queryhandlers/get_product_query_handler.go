package queryhandlers

import (
	"context"
	"product-service/internal/application/queryhandlers/queries"
	"product-service/internal/core/entities"
	"product-service/internal/core/ports"
)

type GetProductQueryHandler struct {
	productRepository ports.ProductRepository
}

func NewGetProductQueryHandler(repo ports.ProductRepository) *GetProductQueryHandler {
	return &GetProductQueryHandler{
		productRepository: repo,
	}
}

func (h *GetProductQueryHandler) Handle(ctx context.Context, query queries.GetProductQuery) (*entities.Product, error) {
	//return h.productRepository.Get(query.ID)
	return nil, nil
}
