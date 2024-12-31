package models

import (
	"fmt"
	"inventory-management/internal/database"
)

func (sm *StockMovement) NewMovement() error {
	var p Product
	if err := database.Db.First(&p, sm.ProductID).Error; err != nil {
		return fmt.Errorf("product doesn't exist")
	}
	if !sm.IsEntry && p.Quantity < sm.Amount {
		return fmt.Errorf("not enough stock for product ID %d", sm.ProductID)
	}
	if sm.IsEntry {
		p.Quantity += sm.Amount
	} else {
		p.Quantity -= sm.Amount
	}
	if err := database.Db.Save(&p).Error; err != nil {
		return fmt.Errorf("internal server error")
	}
	if err := database.Db.Create(&sm).Error; err != nil {
		return fmt.Errorf("internal server error")
	}
	return nil
}

func GetMovementById(id int) (any, error) {
	var sm StockMovement
	if err := database.Db.First(&sm, id).Error; err != nil {
		return nil, fmt.Errorf("product doesn't exist")
	}
	return sm, nil
}