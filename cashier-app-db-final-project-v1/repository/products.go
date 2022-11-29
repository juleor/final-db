package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{db}
}

func (p *ProductRepository) AddProduct(product model.Product) error {
	p.db.First(&product)

	addpro := p.db.Save(&product)
	if addpro != nil {
		return addpro.Error
	}
	return nil // TODO: replace this
}

func (p *ProductRepository) ReadProducts() ([]model.Product, error) {

	return []model.Product{}, nil // TODO: replace this
}

func (p *ProductRepository) DeleteProduct(id uint) error {
	// delpro := p.db.Delete(&ProductRepository{})
	// delpro := p.db.Where("id = ?").Delete(&ProductRepository{})
	delpro := p.db.Delete(&id)
	if delpro.Error != nil {
		return delpro.Error
	}
	return nil // TODO: replace this
}

func (p *ProductRepository) UpdateProduct(id uint, product model.Product) error {
	return nil // TODO: replace this
}
