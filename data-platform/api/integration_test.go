package main

import (
	"api/auth"
	"log"
	"net/http"
)

func TestMain() {
	http.HandleFunc("/signup", auth.SignupHandler)
	http.HandleFunc("/login", auth.LoginHandler)
	http.Handle("/protected", auth.AuthMiddleware(http.HandlerFunc(protectedHandler)))

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected route!"))
}
