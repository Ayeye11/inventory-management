package server

import (
	"fmt"
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
		Handler: Api,
	}
	fmt.Printf("Server listen on %s", server.Addr)
	server.ListenAndServe()
}
