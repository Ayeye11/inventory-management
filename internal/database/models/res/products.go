package res

import (
	"fmt"
	"inventory-management/internal/database"
	"inventory-management/internal/database/models"
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
	case *models.Product:
		return ProductResponse{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Quantity:    v.Quantity,
			Price:       v.Price,
		}
	case []models.Product:
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
	var p []models.Product
	if err := database.Db.Find(&p).Error; err != nil {
		return nil, fmt.Errorf("error")
	}
	return ResProduct(p), nil
}

func GetProductById(id int) (any, error) {
	var p models.Product
	if err := database.Db.Find(&p, id).Error; err != nil {
		return nil, fmt.Errorf("product doesn't exist")
	}
	return ResProduct(p), nil
}
