package main

import (
	"inventory-management/internal/api/server"
	"inventory-management/internal/config"
	"inventory-management/internal/database"
	"inventory-management/internal/database/migrations"
)

func main() {
	config.Env()
	database.InitDatabase()
	migrations.ModelsMigrate()
	server.InitServer()
}
