package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON returns a JSON response to the request
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	if error := json.NewEncoder(w).Encode(data); error != nil {
		log.Fatal(error)
	}
}

// Error returns an error response in JSON format
func Error(w http.ResponseWriter, statusCode int, error error) {
	JSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: error.Error(),
	})
}
