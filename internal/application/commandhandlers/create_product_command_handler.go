package commandhandlers

import (
	"context"
	"product-service/internal/application/commandhandlers/commands"
	"product-service/internal/core/entities"
	"product-service/internal/core/ports"
)

// CreateProductCommandHandler is responsible for handling create product commands.
type CreateProductCommandHandler struct {
	productService ports.ProductService
}

// NewCreateProductCommandHandler creates a new instance of CreateProductCommandHandler.
func NewCreateProductCommandHandler(productService ports.ProductService) *CreateProductCommandHandler {
	return &CreateProductCommandHandler{
		productService: productService,
	}
}

// Handle handles the command to create a new product.
func (h *CreateProductCommandHandler) Handle(ctx context.Context, command commands.CreateProductCommand) error {
	// Create a new product domain entity using the data from the command.
	newProduct := entities.Product{
		Name:        command.Name,
		Description: command.Description,
		Price:       command.Price,
		SKU:         command.SKU,
		Stock:       command.Stock,
		// Set additional fields if necessary
	}

	// Call the domain service to handle business logic and persist the new entity.
	return h.productService.CreateProduct(ctx, &newProduct)
}
