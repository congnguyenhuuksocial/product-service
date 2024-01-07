package productv1

import (
	"context"
	"product-service/internal/core/entities"
	"product-service/internal/core/ports"
)

type productGrpcService struct {
	productService ports.ProductService
	UnimplementedProductServiceServer
}

func NewProductGrpcService(productService ports.ProductService) ProductServiceServer {
	return &productGrpcService{
		productService: productService,
	}
}

func (s *productGrpcService) CreateProduct(ctx context.Context, req *CreateProductRequest) (*ProductResponse, error) {
	product := &entities.Product{
		Name:        req.Name,
		Price:       float64(req.Price),
		Description: req.Description,
	}

	if err := s.productService.CreateProduct(ctx, product); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *productGrpcService) UpdateProduct(ctx context.Context, req *UpdateProductRequest) (*ProductResponse, error) {
	product := &entities.Product{
		Name:        req.Name,
		Price:       float64(req.Price),
		Description: req.Description,
	}

	if err := s.productService.UpdateProduct(ctx, product); err != nil {
		return nil, err
	}

	return nil, nil
}
