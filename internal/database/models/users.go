package models

import (
	"inventory-management/internal/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;size:50;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
}

func (u *User) Add() error {
	if err := database.Db.Create(&u).Error; err != nil {
		return err
	}
	return nil
}
