package commandhandlers

import (
	"context"
	"product-service/internal/application/commandhandlers/commands"
	"product-service/internal/core/entities"
	"product-service/internal/core/ports"
)

type CreateProductCommandHandler struct {
	productService ports.ProductService
}

func NewCreateProductCommandHandler(
	productService ports.ProductService,
) *CreateProductCommandHandler {
	return &CreateProductCommandHandler{
		productService: productService,
	}
}

func (h *CreateProductCommandHandler) Handle(ctx context.Context, command commands.CreateProductCommand) error {
	newProduct := entities.Product{
		Name:        command.Name,
		Description: command.Description,
		Price:       command.Price,
		SKU:         command.SKU,
		Stock:       command.Stock,
	}

	return h.productService.CreateProduct(ctx, &newProduct)
}
