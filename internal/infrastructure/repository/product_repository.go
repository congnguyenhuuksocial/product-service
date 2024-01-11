package repository

import (
	"context"
	"gorm.io/gorm"
	"product-service/internal/core/entities"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(ctx context.Context, product *entities.Product) (*entities.Product, error) {
	return product, r.db.WithContext(ctx).Create(product).Error
}

func (r *ProductRepository) FindByID(ctx context.Context, id string) (*entities.Product, error) {
	var product entities.Product
	err := r.db.WithContext(ctx).First(&product, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) Update(ctx context.Context, product *entities.Product) error {
	return r.db.WithContext(ctx).Save(product).Error
}

func (r *ProductRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&entities.Product{}, "id = ?", id).Error
}
