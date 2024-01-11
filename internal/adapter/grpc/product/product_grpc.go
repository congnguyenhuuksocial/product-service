package product

import (
	"context"
	"net/http"
	messages "product-service/internal/adapter/grpc/product/messages"
	productv1 "product-service/internal/adapter/grpc/product/v1"
	"product-service/internal/application/commandhandlers"
	"product-service/internal/application/commandhandlers/commands"
)

type ProductGrpcService struct {
	productv1.UnimplementedProductServiceServer
	productCreateCommandHandler *commandhandlers.CreateProductCommandHandler
}

func NewProductService(
	productCreateCommandHandler *commandhandlers.CreateProductCommandHandler,
) *ProductGrpcService {
	return &ProductGrpcService{
		productCreateCommandHandler: productCreateCommandHandler,
	}
}

func (s *ProductGrpcService) CreateProduct(ctx context.Context, req *messages.ProductCreateRequest) (*messages.ProductCreateResponse, error) {
	product := commands.CreateProductCommand{
		Name:        req.Name,
		Price:       float64(req.Price),
		Description: req.Description,
	}

	id, err := s.productCreateCommandHandler.Handle(ctx, product)
	if err != nil {
		return nil, err
	}

	return &messages.ProductCreateResponse{
		Message: "success",
		Code:    http.StatusCreated,
		Data: &messages.ProductCreateData{
			Id: id,
		},
	}, nil
}

func (s *ProductGrpcService) UpdateProduct(ctx context.Context, req *messages.ProductUpdateRequest) (*messages.ProductUpdateResponse, error) {
	return nil, nil
}

func (s *ProductGrpcService) DeleteProduct(ctx context.Context, req *messages.ProductDeleteRequest) (*messages.ProductDeleteResponse, error) {
	return nil, nil
}

func (s *ProductGrpcService) GetProductById(ctx context.Context, req *messages.ProductGetRequest) (*messages.ProductGetResponse, error) {
	return nil, nil
}

func (s *ProductGrpcService) GetProducts(ctx context.Context, req *messages.ProductListRequest) (*messages.ProductListResponse, error) {
	return nil, nil
}
