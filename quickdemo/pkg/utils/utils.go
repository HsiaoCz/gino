package utils

import (
	"encoding/json"
	"net/http"
)

type H map[string]interface{}

func ResponseJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

func ResponseString(w http.ResponseWriter, code int, str string) {
	w.WriteHeader(code)
	w.Write([]byte(str))
}
