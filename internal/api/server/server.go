package server

import (
	"fmt"
	"inventory-management/internal/api/router/middlewares"
	"net/http"
	"os"
)

var Api = http.NewServeMux()

func InitServer() {
	addr := os.Getenv("SERVER_PORT")
	if addr == "" {
		addr = ":8080"
	}
	server := http.Server{
		Addr:    addr,
		Handler: middlewares.ChainMiddlewares(Api, middlewares.Middlewares...),
	}
	fmt.Printf("Server listen on %s\n", server.Addr)
	server.ListenAndServe()
}
