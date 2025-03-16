package handlers

import (
	"encoding/json"
	"net/http"
	"project/model"
)

func HandleNations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(model.Nations)
	w.Write([]byte(data))
}
