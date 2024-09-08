package auth

import (
	"bytes"
	"encoding/json"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoginHandler(t *testing.T) {
	tests := []struct {
		name               string
		email              string
		password           string
		expectedStatusCode int
		expectedBody       string
	}{
		{
			name:               "Valid credentials",
			email:              "bob",
			password:           "1234",
			expectedStatusCode: http.StatusOK,
			expectedBody:       "Login successful!",
		},
		{
			name:               "Invalid credentials",
			email:              "bob",
			password:           "wrongpassword",
			expectedStatusCode: http.StatusUnauthorized,
			expectedBody:       "Invalid credentials",
		},
		{
			name:               "Bad request",
			email:              "",
			password:           "",
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       "Bad Request",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			creds := Credentials{Email: tt.email, Password: tt.password}
			body, _ := json.Marshal(creds)

			req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
			if err != nil {
				t.Fatal("Error creating request:", err)
			}
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			LoginHandler(rr, req)

			if strings.TrimSpace(rr.Body.String()) != tt.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v",
					strings.TrimSpace(rr.Body.String()), tt.expectedBody)
			}

			if status := rr.Code; status != tt.expectedStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatusCode)
			}
		})
	}
}

func TestSignUpHandler(t *testing.T) {
	tests := []struct {
		name               string
		email              string
		password           string
		expectedStatusCode int
		expectedBody       string
	}{
		{
			name:               "Successful signup",
			email:              "newuser@example.com",
			password:           "password123",
			expectedStatusCode: http.StatusCreated,
			expectedBody:       "Signup successful",
		},
		{
			name:               "Email exists",
			email:              "bob",
			password:           "newpassword",
			expectedStatusCode: http.StatusConflict,
			expectedBody:       "Email exists or password doesn't match requirements",
		},
		{
			name:               "Bad request",
			email:              "",
			password:           "",
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       "Email and password are required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			creds := Credentials{Email: tt.email, Password: tt.password}
			body, _ := json.Marshal(creds)

			req, err := http.NewRequest("POST", "/signUp", bytes.NewBuffer(body))
			if err != nil {
				t.Fatal("Error creating request:", err)
			}
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			SignupHandler(rr, req)

			if strings.TrimSpace(rr.Body.String()) != tt.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v",
					strings.TrimSpace(rr.Body.String()), tt.expectedBody)
			}

			if status := rr.Code; status != tt.expectedStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatusCode)
			}
		})
	}
}
