package handlers

import (
	"fmt"
	"inventory-management/internal/database/models"
	"inventory-management/pkg/utils"
	"net/http"
	"strconv"
)

func GetMovements(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Page")
}

func GetMovementsId(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
	sm, err := models.GetMovementById(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
	utils.WriteJson(w, http.StatusOK, sm)
}

func PostMovements(w http.ResponseWriter, r *http.Request) {
	var sm models.StockMovement
	if err := utils.ParseJson(r, &sm); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if err := sm.NewMovement(); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteMessage(w, http.StatusOK, "movement added successfully")
}
