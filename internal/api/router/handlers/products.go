package handlers

import (
	"inventory-management/internal/database/models"
	"inventory-management/internal/database/models/res"
	"inventory-management/pkg/utils"
	"net/http"
	"strconv"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	data, err := res.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
	}
	utils.WriteJson(w, http.StatusOK, data)
}

func GetProductsId(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
	data, err := res.GetProductById(id)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
	}
	utils.WriteJson(w, http.StatusOK, data)
}

func PostProducts(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	if err := utils.ParseJson(r, &p); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if err := p.Add(); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteMessage(w, http.StatusCreated, "product created successfully")
}

func PutProductsId(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
	if err := utils.ParseJson(r, &p); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if err := p.Update(id); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteMessage(w, http.StatusCreated, "product updated successfully")
}

func DeleteProductsId(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
	if err := p.Delete(id); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteMessage(w, http.StatusCreated, "product deleted successfully")
}
