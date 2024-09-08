package routes

import (
	"api/internal/handlers"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"golang.org/x/time/rate"
)

var uploadedFilePath string
var filenameS3 string = ""
var limiter *rate.Limiter

type Options struct {
	method   string
	filename string
	bucket   string
	region   string
	expiry   time.Duration
}

type Lab struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Location  string `json:"location"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Routes initializes and returns the router with all routes.
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)          // Logs the start and end of each request with the elapsed processing time
	router.Use(middleware.Recoverer)       // Recovers from panics and writes a 500
	router.Use(rateLimitMiddleware)        //Uses the limiter
	router.Route("/", func(r chi.Router) { //example
		r.Get("/", indexHandler)
	})
	router.Route("/videoThumbnail", func(r chi.Router) {
		r.Get("/", videoThumbnailHandler)
	})
	router.Route("/videoUpload", func(r chi.Router) {
		r.Post("/", videoUploadHandler)
	})
	router.Route("/generate-presigned-url", func(r chi.Router) {
		r.Post("/", generatePresignedURLHandler)
	})
	router.Route("/labs", func(r chi.Router) {
		r.Get("/", handlers.GetLabsHandler)
	})
	router.Route("/users", func(r chi.Router) {
		r.Get("/", handlers.GetUsersHandler)
	})
	router.Route("/devices", func(r chi.Router) {
		r.Get("/", handlers.GetDevicesHandler)
	})
	return router
}
func initializeRateLimiter() { //one of the first methods that is run in sequntial order
	godotenv.Load("config/.env")
	rateLimitRequests, err := strconv.Atoi(os.Getenv("RATE_LIMIT_REQUESTS"))
	if err != nil {
		rateLimitRequests = 5 //here in case it doesn't successful retrive the value
	}
	rateLimitSeconds, err := strconv.Atoi(os.Getenv("RATE_LIMIT_DURATION"))
	if err != nil {
		rateLimitSeconds = 180 //here in case it doesn't successful retrive the value
	}
	rateLimitDuration := time.Duration(rateLimitSeconds) * time.Second
	limiter = rate.NewLimiter(rate.Every(rateLimitDuration), rateLimitRequests)
}
func init() {
	initializeRateLimiter()
}
func rateLimitMiddleware(next http.Handler) http.Handler { //checks to see if the request made is allowed by the rate limiter
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r) //passes the call to the next function in the chain
	})
}
func generatePresignedURLHandler(w http.ResponseWriter, r *http.Request) {
	var reqData struct {
		Filename string `json:"filename"`
	}
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	extension := strings.ToLower(filepath.Ext(reqData.Filename))
	var prefix string

	switch extension {
	case ".jpg", ".jpeg", ".png", ".gif":
		prefix = "images/"
	case ".pdf":
		prefix = "documents/"
	case ".mp4", ".avi", ".mov":
		prefix = "videos/"
	default:
		prefix = "others/"
	}

	filenameWithPrefix := prefix + reqData.Filename

	duration, _ := strconv.Atoi(os.Getenv("DURATION"))
	options := Options{
		method:   "PUT",
		filename: filenameWithPrefix,
		bucket:   os.Getenv("BUCKET_NAME"),
		region:   os.Getenv("AWS_REGION"),
		expiry:   time.Minute * time.Duration(duration),
	}
	urlStr, err := generatePresignedURL(options)
	if err != nil {
		log.Println("Error presigning request:", err)
		http.Error(w, "Failed to sign request", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"url": urlStr,
	})
}
func generatePresignedURL(options Options) (string, error) {
	service_Account, account_Number := os.Getenv("SERVICE_ACCOUNT_NAME"), os.Getenv("ACCOUNT_NUMBER")
	fmt.Sprintf("arn:aws:iam::%s:/user/%s", service_Account, account_Number)
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(options.region),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			""),
	})
	if err != nil {
		log.Println("Error creating AWS session:", err)
		return "", err
	}
	svc := s3.New(sess)
	var req *request.Request
	switch options.method {
	case "PUT":
		req, _ = svc.PutObjectRequest(&s3.PutObjectInput{
			Bucket: aws.String(options.bucket),
			Key:    aws.String(options.filename),
		})
	case "GET":
		req, _ = svc.GetObjectRequest(&s3.GetObjectInput{
			Bucket: aws.String(options.bucket),
			Key:    aws.String(options.filename),
		})
	case "DELETE":
		req, _ = svc.DeleteObjectRequest(&s3.DeleteObjectInput{
			Bucket: aws.String(options.bucket),
			Key:    aws.String(options.filename),
		})
	default:
		return "", fmt.Errorf("not a valid method")
	}
	filenameS3 = options.filename
	urlStr, err := req.Presign(options.expiry)
	if err != nil {
		log.Println("Error presigning request:", err)
		return "", err
	}
	return urlStr, nil
}
func indexHandler(w http.ResponseWriter, r *http.Request) { // study what exactly these methods do
	tmpl, err := template.ParseFiles("public/index.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
func videoThumbnailHandler(w http.ResponseWriter, r *http.Request) {
	if uploadedFilePath != "" {
		http.ServeFile(w, r, uploadedFilePath)
	} else {
		http.NotFound(w, r)
	}
}
func videoUploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Error retrieving file", http.StatusInternalServerError)
			return
		}
		defer file.Close()
		filename := fileHeader.Filename
		savePath := filepath.Join("videoUploads", filename)
		savedFile, err := os.Create(savePath)
		if err != nil {
			http.Error(w, "Error creating file", http.StatusInternalServerError)
			return
		}
		defer savedFile.Close()
		_, err = io.Copy(savedFile, file)
		if err != nil {
			http.Error(w, "Error writing file", http.StatusInternalServerError)
			return
		}
		uploadedFilePath = savedFile.Name()
		log.Println("Video uploaded successfully")
	}
}
