package main

import (
	"inventory-management/internal/api/router"
	"inventory-management/internal/api/server"
	"inventory-management/internal/config"
	"inventory-management/internal/database"
)

func main() {
	config.Env()
	database.InitDatabase()
	database.Migrate()
	router.InitRouter()
	server.InitServer()
}
