package middlewares

import (
	"inventory-management/internal/api/auth"
	"inventory-management/pkg/utils"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.GetToken(r)
		if err != nil {
			utils.WriteError(w, http.StatusUnauthorized, err)
			return
		}
		valid, err := auth.VerifyToken(token)
		if err != nil || !valid {
			utils.WriteError(w, http.StatusUnauthorized, err)
			return
		}
		next.ServeHTTP(w, r)
	})
}