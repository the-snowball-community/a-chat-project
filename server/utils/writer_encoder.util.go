package utils

import (
	"encoding/json"
	"net/http"
)

func WriteEncoder(w http.ResponseWriter, data any) {
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}
}
