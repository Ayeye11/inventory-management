package handlers

import (
	"inventory-management/internal/api/auth"
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
		return
	}
	hash, err := auth.HashPassword(u.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	u.Password = hash
	if err := u.Add(); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteMessage(w, http.StatusOK, "user created successfully")
}

// login
func GetLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login page"))
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	var u models.User
	if err := utils.ParseJson(r, &u); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if !auth.ValidateUser(u) {
		utils.WriteMessage(w, http.StatusUnauthorized, "incorrect username or password")
		return
	}
	token, err := auth.GenerateJWT(u)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.SendToken(w, http.StatusOK, token)
}
