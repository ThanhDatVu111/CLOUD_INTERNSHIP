package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var jwtKey = []byte("my_secret_key")
var jwtExpiryDuration time.Duration

// User represents a user in the system
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Credentials represents a struct to hold Email and password
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Claims represents a struct for JWT claims
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// mock users for testing
var users = map[string]string{
	"testuser": "password123",
	"bob":      "1234",
	"a":        "b",
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default JWT expiry duration of 1 minute")
	}
}

// SignupHandler handles user signup
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Check if the Email already exists
	if _, ok := users[creds.Email]; ok {
		http.Error(w, "Email exists or password doesn't match requirements", http.StatusConflict)
		return
	}
	if creds.Email == "" || creds.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	// Add the new user to the mock database
	users[creds.Email] = creds.Password

	// Return success response
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Signup successful")) // Write response body
	log.Println("Signup successful for email:", creds.Email)
}

// LoginHandler handles user login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if creds.Email == "" || creds.Password == "" {
		log.Println("Empty email or password")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[creds.Email]
	if !ok || expectedPassword != creds.Password {
		log.Println("Invalid credentials for email:", creds.Email)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(jwtExpiryDuration)
	claims := &Claims{
		Email: creds.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println("Error signing JWT:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful!")) // Write response body
	log.Println("Login successful for email:", creds.Email)
}

// AuthMiddleware validates the JWT token
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		tknStr := c.Value
		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

func CsrfToken() {

}
