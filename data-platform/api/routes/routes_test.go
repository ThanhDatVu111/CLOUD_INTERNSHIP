package routes

import (
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	"net/http"
	"net/http/httptest"

	"github.com/joho/godotenv"
)

func TestRateLimitMiddleware(t *testing.T) {
	godotenv.Load()
	os.Setenv("RATE_LIMIT_REQUESTS", "2") //sets the value of the .env values
	os.Setenv("RATE_LIMIT_DURATION", "1s")
	initializeRateLimiter() //tests the rate limiter set up in the routes.go

	var numTests, _ = strconv.Atoi(os.Getenv("RATE_LIMIT_REQUESTS"))

	// Dummy handler to test the middleware
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Wrap the handler with the rate limiter middleware
	limitedHandler := rateLimitMiddleware(handler)

	// Create a test request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	for i := 0; i < numTests; i++ {
		rr := httptest.NewRecorder() //records the response messages
		limitedHandler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("expected status OK, got %v", status) //making sure we aren't having an error here, we're still supposed to be within the number of allowed requests
		}
	}

	// Third request should exceed the rate limit
	rr := httptest.NewRecorder()
	limitedHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusTooManyRequests {
		t.Errorf("expected status Too Many Requests, got %v", status)
	}
}

func TestGeneratePresignedURL_PUT(t *testing.T) {
	godotenv.Load()
	options := Options{
		method:   "PUT",
		filename: "Test.txt",
		bucket:   os.Getenv("BUCKET_NAME"),
		region:   os.Getenv("AWS_REGION"),
		expiry:   time.Duration(15),
	}
	URL, err := generatePresignedURL(options)
	if err != nil {
		t.Fatalf("Error creating AWS presigned link: %v", err)
	} else {
		log.Println(URL)
	}
}

func TestGeneratePresignedURL_GET(t *testing.T) {
	godotenv.Load()
	options := Options{
		method:   "GET",
		filename: "puppy.png",
		bucket:   os.Getenv("BUCKET_NAME"),
		region:   os.Getenv("AWS_REGION"),
		expiry:   time.Duration(15),
	}
	URL, err := generatePresignedURL(options)
	if err != nil {
		t.Fatalf("Error creating AWS presigned link for a GET request: %v", err)
	} else {
		log.Println(URL)
	}
}

func TestGeneratePresignedURL_DELETE(t *testing.T) {
	godotenv.Load()
	options := Options{
		method:   "DELETE",
		filename: "puppy.png",
		bucket:   os.Getenv("BUCKET_NAME"),
		region:   os.Getenv("AWS_REGION"),
		expiry:   time.Duration(15),
	}
	URL, err := generatePresignedURL(options)
	if err != nil {
		t.Fatalf("Error creating AWS presigned link for a GET request: %v", err)
	} else {
		log.Println(URL)
	}
}
