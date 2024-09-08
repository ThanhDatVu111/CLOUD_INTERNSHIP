package main

import (
	"log"
	"net/http"
	platform "portal/platform/src/data"

	"github.com/rs/cors"
)

func main() {

	r := platform.Routes()

	// CORS configuration, comment this out if this causes any error
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // platform frontend
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	})

	// CORS middleware handler
	handler := c.Handler(r)

	http.Handle("/", handler)
	log.Println("Server started at localhost")
	log.Fatal(http.ListenAndServe(":5174", nil))
}
