package handlers

import (
	"inventory-management/internal/database/models"
	"inventory-management/pkg/utils"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	data, err := models.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
	}
	utils.WriteJson(w, http.StatusOK, data)
}
