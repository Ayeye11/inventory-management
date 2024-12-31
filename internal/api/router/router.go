package router

import (
	"inventory-management/internal/api/router/handlers"
	"inventory-management/internal/api/router/urls"
	"inventory-management/internal/api/server"
)

func InitRouter() {
	home()
	inventory()
	products()
	StockMovement()
	notFound()
}

//handlers

func home() {
	server.Api.HandleFunc("GET "+urls.Home, handlers.GetHome)
}

func inventory() {
	server.Api.HandleFunc("POST "+urls.Inventory, handlers.GetInventory)
}

func products() {
	server.Api.HandleFunc("GET "+urls.Product, handlers.GetProducts)
	server.Api.HandleFunc("GET "+urls.Product_paramId, handlers.GetProductsId)
	server.Api.HandleFunc("POST "+urls.Product, handlers.PostProducts)
	server.Api.HandleFunc("PUT "+urls.Product_paramId, handlers.PutProductsId)
	server.Api.HandleFunc("DELETE "+urls.Product_paramId, handlers.DeleteProductsId)
}

func StockMovement() {
	server.Api.HandleFunc("GET "+urls.StockMovement, handlers.GetMovements)
	server.Api.HandleFunc("GET "+urls.StockMovement__paramId, handlers.GetMovementsId)
	server.Api.HandleFunc("POST "+urls.StockMovement, handlers.PostMovements)
}

func notFound() {
	server.Api.HandleFunc("/", handlers.UseNotFound)
}
