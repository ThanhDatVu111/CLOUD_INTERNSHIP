package handlers

import (
	"api/internal/database/postgres"
	"api/internal/models"
	"encoding/json"
	"log"
	"net/http"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetAllUsers(postgres.NewDB(), 100)
	if err != nil {
		http.Error(w, "Error getting user data", http.StatusInternalServerError)
		return
	}
	log.Println(users)
	json, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Error comverting users to JSON format", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
