package models

import (
	"fmt"
	"inventory-management/internal/database"
)

func (p *Product) Add() error {
	p.Quantity = 0
	if err := database.Db.Create(&p).Error; err != nil {
		return fmt.Errorf("this product already exists")
	}
	return nil
}

func (p *Product) Update(id int) error {
	var product Product
	if err := database.Db.First(&product, id).Error; err != nil {
		return fmt.Errorf("this product doesn't exist")
	}
	if err := database.Db.Model(&product).Updates(p).Error; err != nil {
		return fmt.Errorf("failed to update product")
	}
	return nil
}

func (p *Product) Delete(id int) error {
	if err := database.Db.First(&p, id).Error; err != nil {
		return fmt.Errorf("this product doesn't exist")
	}
	if err := database.Db.Delete(&p).Error; err != nil {
		return fmt.Errorf("failed to delete product")
	}
	return nil
}
