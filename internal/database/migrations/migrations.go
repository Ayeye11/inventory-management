package migrations

import (
	"inventory-management/internal/database"
	"inventory-management/internal/database/models"
	"log"
)

func Migrate() {
	err := database.Db.AutoMigrate(&models.Product{}, &models.StockMovement{})
	if err != nil {
		log.Fatal(err)
	}
}
