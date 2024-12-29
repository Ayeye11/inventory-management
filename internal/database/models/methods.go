package models

import (
	"fmt"
	"inventory-management/internal/database"
)

func AddProduct(v Product) error {
	if err := database.Db.Create(&v).Error; err != nil {
		return fmt.Errorf("this product already exists")
	}
	return nil
}

func UpdateProduct(v Product, id int) error {
	var product Product
	if err := database.Db.First(&product, id).Error; err != nil {
		return fmt.Errorf("this product doesn't exist")
	}
	if err := database.Db.Model(&product).Updates(v).Error; err != nil {
		return fmt.Errorf("failed to update product")
	}
	return nil
}

func DeleteProduct(id int) error {
	var product Product
	if err := database.Db.First(&product, id).Error; err != nil {
		return fmt.Errorf("this product doesn't exist")
	}
	if err := database.Db.Delete(&product).Error; err != nil {
		return fmt.Errorf("failed to delete product")
	}
	return nil
}

func (a *Product) NewMovement(v StockMovement) error {
	if err := database.Db.First(a, v.ProductID).Error; err != nil {
		return fmt.Errorf("product doesn't exist")
	}
	if !v.IsEntry {
		if a.Quantity < v.Amount {
			return fmt.Errorf("not enough stock")
		}
		a.Quantity -= v.Amount
	} else {
		a.Quantity += v.Amount
	}
	if err := database.Db.Save(a).Error; err != nil {
		return err
	}
	if err := database.Db.Create(&v).Error; err != nil {
		return err
	}
	return nil
}
