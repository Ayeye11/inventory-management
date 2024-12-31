package handlers

import (
	"inventory-management/internal/database/models/res"
	"inventory-management/pkg/utils"
	"net/http"
)

func GetInventory(w http.ResponseWriter, r *http.Request) {
	data, err := res.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
	}
	utils.WriteJson(w, http.StatusOK, data)
}
