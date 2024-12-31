package handlers

import (
	"fmt"
	"net/http"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "Home page")
}
