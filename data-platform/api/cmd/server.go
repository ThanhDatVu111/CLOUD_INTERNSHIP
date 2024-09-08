package main

import (
	"log"
	"net/http"
	"os"

	"api/routes"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	godotenv.Load("config/.env", "internal/database/postgres/.env")
	directories := []string{"videoUploads"}

	for _, dir := range directories {
		if _, err := os.Stat(dir); err != nil {
			if os.IsNotExist(err) {
				log.Fatalf("Failed to find the %s directory: %v\n", dir, err)
			} else {
				log.Fatalf("Error checking the %s directory: %v\n", dir, err)
			}
		}
	}

	router := routes.Routes()

	// CORS configuration, comment this out if this causes any error
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081"}, //my Expo app url
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	})

	// CORS middleware handler
	handler := c.Handler(router)

	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css", fs))

	http.Handle("/", handler)
	// todo GetJwtDuration to expire the JWT token
	log.Println("Listening on port :3000")
	http.ListenAndServe(":3000", nil)
}
