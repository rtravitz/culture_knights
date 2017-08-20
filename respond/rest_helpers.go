package respond

import (
	"encoding/json"
	"net/http"
)

func WithError(w http.ResponseWriter, code int, message string) {
	WithJSON(w, code, map[string]string{"error": message})
}

func WithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(response)
}
