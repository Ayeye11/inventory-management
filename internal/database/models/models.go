package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `gorm:"size:255;not null"`
	Description string  `gorm:"size:500;not null"`
	Price       float64 `gorm:"not null"`
	Quantity    int     `gorm:"not null;default:0"`
	CategoryID  int     `gorm:"not null"`
	SupplierID  int     `gorm:"not null"`

	Category Category `gorm:"foreignKey:CategoryID;references:ID"`
	Supplier Supplier `gorm:"foreignKey:SupplierID;references:ID"`
}

type Category struct {
	gorm.Model
	Name string `gorm:"size:255;not null"`
}

type Supplier struct {
	gorm.Model
	Name    string `gorm:"size:255;not null"`
	Contact string `gorm:"size:255"`
}

type StockMovement struct {
	gorm.Model
	Type      string    `gorm:"size:10;not null"`
	Amount    int       `gorm:"not null"`
	Date      time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	ProductID uint      `gorm:"not null"`

	Product Product `gorm:"foreignKey:ProductID;references:ID"`
}
