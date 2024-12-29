package migrations

import (
	"inventory-management/internal/database"
	"inventory-management/internal/database/models"
	"log"
)

func ModelsMigrate() {
	err := database.Db.AutoMigrate(&models.Product{}, &models.Category{}, &models.Supplier{}, &models.StockMovement{})
	if err != nil {
		log.Fatal(err)
	}
}
