package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJson(r *http.Request, target any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	if err := json.NewDecoder(r.Body).Decode(&target); err != nil {
		return err
	}
	return nil
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) error {
	return WriteJson(w, status, map[string]string{"error": err.Error()})
}

func WriteMessage(w http.ResponseWriter, status int, message string) error {
	return WriteJson(w, status, map[string]string{"message": message})
}
