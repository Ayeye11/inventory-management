package models

import (
	"fmt"
	"inventory-management/internal/database"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `gorm:"unique;size:255;not null"`
	Description string  `gorm:"size:500;not null"`
	Price       float64 `gorm:"not null"`
	Quantity    int     `gorm:"not null;default:0"`
}

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
