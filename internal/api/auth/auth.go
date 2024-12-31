package auth

import (
	"inventory-management/internal/database"
	"inventory-management/internal/database/models"
)

func ValidateUser(u models.User) bool {
	var userDb models.User
	if err := database.Db.First(&userDb, "Username=?", u.Username).Error; err != nil {
		return false
	}
	return CheckPassword(userDb.Password, u.Password)
}
