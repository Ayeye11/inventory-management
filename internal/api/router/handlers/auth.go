package handlers

import (
	"inventory-management/internal/database/models"
	"inventory-management/pkg/utils"
	"net/http"
)

// register
func GetRegister(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Register page"))
}

func PostRegister(w http.ResponseWriter, r *http.Request) {
	var u models.User
	if err := utils.ParseJson(r, &u); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

}

// login
func GetLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login page"))
}
