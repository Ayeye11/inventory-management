package database

import (
	"inventory-management/internal/database/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDatabase() {
	dsn := os.Getenv("DSN")
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func Migrate() {
	err := Db.AutoMigrate(&models.Product{}, &models.Category{}, &models.Supplier{}, &models.StockMovement{})
	if err != nil {
		log.Fatal(err)
	}
}
