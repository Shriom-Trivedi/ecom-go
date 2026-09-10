package json

import (
	"encoding/json"
	"net/http"
)

func Write(w http.ResponseWriter, status int, data any) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func Read(r *http.Request, data any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Do not accept unwanted parameters from Body, if found, returns an error.
	
	return decoder.Decode(data)
}
