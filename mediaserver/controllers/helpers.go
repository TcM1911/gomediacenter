package controllers

import (
	"encoding/json"
	"net/http"
)

func writeJsonBody(w http.ResponseWriter, v interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
