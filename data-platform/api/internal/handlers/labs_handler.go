package handlers

import (
	"api/internal/database/postgres"
	"api/internal/models"
	"encoding/json"
	"log"
	"net/http"
)


func GetLabsHandler(w http.ResponseWriter, r *http.Request) {
	labs, err := models.GetAllLabs(postgres.NewDB())
	if err != nil {
		http.Error(w, "Error getting lab data", http.StatusInternalServerError)
		return
	}
	log.Println(labs)
	json, err := json.Marshal(labs)
	if err != nil {
		http.Error(w, "Error comverting labs to JSON format", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}