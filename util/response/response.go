package util

import (
	"encoding/json"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, status int, success int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	res := map[string]interface{}{
		"success": success,
		"data":    data,
		"message": message,
	}
	json.NewEncoder(w).Encode(res)
}
