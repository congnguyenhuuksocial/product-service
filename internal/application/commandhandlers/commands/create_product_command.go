package commands

// CreateProductCommand carries the data required to create a product.
type CreateProductCommand struct {
	Name        string
	Description string
	Price       float64
	SKU         string
	Stock       int
}
