package router

import (
	"inventory-management/internal/api/router/handlers"
	"inventory-management/internal/api/server"
)

func InitRouter() {
	//404
	notFound()
}

//products

// notFound
func notFound() {
	server.Api.HandleFunc("/", handlers.UseNotFound)
}
