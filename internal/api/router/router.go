package router

import (
	"inventory-management/internal/api/router/handlers"
	"inventory-management/internal/api/router/urls"
	"inventory-management/internal/api/server"
)

func InitRouter() {
	home()
	login()
	inventory()
	products()
	StockMovement()
	notFound()
}

//public

func home() {
	server.Api.HandleFunc("GET "+urls.Home, handlers.GetHome)
}

func login() {
	server.Api.HandleFunc("GET "+urls.Auth_register, handlers.GetRegister)
	server.Api.HandleFunc("POST "+urls.Auth_register, handlers.PostRegister)
	server.Api.HandleFunc("GET "+urls.Auth_login, handlers.GetLogin)
	server.Api.HandleFunc("POST "+urls.Auth_login, handlers.PostLogin)
}

//private

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
