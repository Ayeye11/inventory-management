package middlewares

import (
	"fmt"
	"net/http"
	"path"
)

type Middleware func(http.Handler) http.Handler

var Middlewares = []Middleware{
	redirectHome,
	logRequest,
}

func ChainMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if path.Ext(r.URL.Path) != ".ico" {
			fmt.Printf("Method: %s, URL: %s\n", r.Method, r.URL)
		}
		next.ServeHTTP(w, r)
	})
}

func redirectHome(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/home", http.StatusFound)
		}
		next.ServeHTTP(w, r)
	})
}
