package handlers

import (
    "api/internal/database/postgres"
	"api/internal/models"
	"encoding/json"
	"log"
	"net/http"
)

// GetDevicesHandler handles the /devices endpoint
func GetDevicesHandler(w http.ResponseWriter, r *http.Request) {
    devices, err := models.GetAllDevices(postgres.NewDB(), 100)
    if err != nil {
		http.Error(w, "Error getting device data", http.StatusInternalServerError)
		return
	}
    log.Println(devices)
	json, err := json.Marshal(devices)
	if err != nil {
		http.Error(w, "Error comverting devices to JSON format", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}