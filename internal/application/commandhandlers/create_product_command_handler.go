package commandhandlers

import (
	"context"
	"github.com/go-playground/validator/v10"
	"product-service/internal/application/commandhandlers/commands"
	"product-service/internal/core/entities"
	"product-service/internal/core/ports"
)

type CreateProductCommandHandler struct {
	productService ports.ProductService
	validate       *validator.Validate
}

func NewCreateProductCommandHandler(
	productService ports.ProductService,
	validate *validator.Validate,
) *CreateProductCommandHandler {
	return &CreateProductCommandHandler{
		productService: productService,
		validate:       validate,
	}
}

func (h *CreateProductCommandHandler) Handle(ctx context.Context, command commands.CreateProductCommand) (uint32, error) {
	err := h.validate.Struct(command)
	if err != nil {
		return 0, err
	}
	newProduct := entities.Product{
		Name:        command.Name,
		Description: command.Description,
		Price:       command.Price,
		SKU:         command.SKU,
		Stock:       int64(command.Stock),
	}

	product, err := h.productService.CreateProduct(ctx, &newProduct)

	return uint32(product.ID), err
}
