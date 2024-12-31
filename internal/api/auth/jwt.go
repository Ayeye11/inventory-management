package auth

import (
	"fmt"
	"inventory-management/internal/database/models"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey string = os.Getenv("TOKEN_KEY")

func GenerateJWT(u models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["username"] = u.Username
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func VerifyToken(userToken string) (bool, error) {
	token, err := jwt.Parse(userToken, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("error")
		}
		return secretKey, nil
	})
	if err != nil {
		return false, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				return false, fmt.Errorf("token expired")
			}
		}
		return true, nil
	}
	return false, fmt.Errorf("invalid token")
}

func GetToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("missing token")
	}
	return strings.TrimPrefix(authHeader, "Bearer "), nil
}
