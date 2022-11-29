package repository

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return CartRepository{db}
}

func (c *CartRepository) ReadCart() ([]model.JoinCart, error) {
	c.db.First([]model.JoinCart{})

	return []model.JoinCart{}, nil // TODO: replace this
}

func (c *CartRepository) AddCart(product model.Product) error {

	// addcar := c.db.Save(&product)
	// if addcar != nil {
	// 	return addcar.Error
	// }

	if product.Stock == 0 {
		return fmt.Errorf("Out of Stock")
	}

	hargacart := product.Price
	if product.Discount != 0 {
		hargacart = product.Price - ((product.Discount / 100) * product.Price)
	}

	cart := model.Cart{
		ProductID:  product.ID,
		Quantity:   1,
		TotalPrice: hargacart,
	}
	if err := c.db.Where("product_id=? AND deleted_at IS NULL", product.ID).First(&cart).Error; err != nil {
		c.db.Transaction(func(tx *gorm.DB) error {
			product.Stock = product.Stock - 1
			if err := tx.Table("product").Where("id = ? AND deleted_at IS NULL", product.ID).Updates(product).Error; err != nil {
				return err
			}
			if err := tx.Create(&cart).Error; err != nil {
				return err
			}
			return nil

		})

		// TODO: replace this
	}
	return nil
}

func (c *CartRepository) DeleteCart(id uint, productID uint) error {

	delcar := c.db.Delete(&id, &productID)
	if delcar.Error != nil {
		return delcar.Error
	}
	return nil // TODO: replace this
}

func (c *CartRepository) UpdateCart(id uint, cart model.Cart) error {
	return nil // TODO: replace this
}
