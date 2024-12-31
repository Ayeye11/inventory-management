package models

import (
	"fmt"
	"inventory-management/internal/database"
)

type ProductResponse struct {
	ID          uint
	Name        string
	Description string
	Quantity    int
	Price       float64
}

func ResProduct(a any) any {
	switch v := a.(type) {
	case *Product:
		return ProductResponse{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Quantity:    v.Quantity,
			Price:       v.Price,
		}
	case []Product:
		var response []ProductResponse
		for _, item := range v {
			response = append(response, ProductResponse{
				ID:          item.ID,
				Name:        item.Name,
				Description: item.Description,
				Quantity:    item.Quantity,
				Price:       item.Price,
			})
		}
		return response
	default:
		return nil
	}
}

func GetProducts() (any, error) {
	var p Product
	if err := database.Db.Find(&p).Error; err != nil {
		return nil, fmt.Errorf("no products to show")
	}
	return ResProduct(p), nil
}
