package commands

// CreateProductCommand carries the data required to create a product.
type CreateProductCommand struct {
	Name        string  `validate:"required"`
	Description string  `validate:"required;max=1000"`
	Price       float64 `validate:"required;min=0"`
	SKU         string  `validate:"required"`
	Stock       int     `validate:"required"`
}
