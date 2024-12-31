package res

import (
	"fmt"
	"inventory-management/internal/database"
	"inventory-management/internal/database/models"
)

func GetMovementById(id int) (any, error) {
	var sm models.StockMovement
	if err := database.Db.First(&sm, id).Error; err != nil {
		return nil, fmt.Errorf("product doesn't exist")
	}
	return sm, nil
}
