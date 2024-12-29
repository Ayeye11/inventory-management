package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `gorm:"unique;size:255;not null"`
	Description string  `gorm:"size:500;not null"`
	Price       float64 `gorm:"not null"`
	Quantity    int     `gorm:"not null;default:0"`
}

type StockMovement struct {
	gorm.Model
	IsEntry   bool      `gorm:"not null"`
	Amount    int       `gorm:"not null"`
	Date      time.Time `gorm:"type:date;not null;default:CURRENT_DATE"`
	ProductID uint      `gorm:"not null"`

	Product Product `gorm:"foreignKey:ProductID;references:ID"`
}
