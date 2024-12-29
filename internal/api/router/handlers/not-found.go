package handlers

import (
	"fmt"
	"net/http"
)

func UseNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "Page not found")
}
