package productv1

import (
	"context"
	"product-service/internal/core/entities"
	"product-service/internal/core/ports"
)

type ProductGrpcService struct {
	productService ports.ProductService
	UnimplementedProductServiceServer
}

func NewProductService(productService ports.ProductService) *ProductGrpcService {
	return &ProductGrpcService{
		productService: productService,
	}
}

func (s *ProductGrpcService) CreateProduct(ctx context.Context, req *CreateProductRequest) (*ProductResponse, error) {
	product := &entities.Product{
		Name:        req.Name,
		Price:       float64(req.Price),
		Description: req.Description,
	}

	if err := s.productService.CreateProduct(ctx, product); err != nil {
		return nil, err
	}

	return &ProductResponse{
		Id:          0,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}, nil
}

func (s *ProductGrpcService) UpdateProduct(ctx context.Context, req *UpdateProductRequest) (*ProductResponse, error) {
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
