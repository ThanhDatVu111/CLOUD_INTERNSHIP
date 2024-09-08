package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// checkErr logs the error with additional context and terminates the program for critical errors.
func checkErr(err error, context string) {
	if err != nil {
		log.Fatalf("Context: %s, Error: %v", context, err)
	}
}

func init() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func GetConnectionHandler() (*sql.DB, error) {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	checkErr(err, "DB_PORT is not set or is empty")
	user := os.Getenv("DB_USER")
	pw := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("port=%d user=%s password=%s dbname=%s sslmode=disable", port, user, pw, db_name)
	fmt.Println("Connection String:", connStr)

	db, err := sql.Open("postgres", connStr)
	checkErr(err, "opening database connection")

	err = db.Ping()
	checkErr(err, "pinging database")

	return db, nil
}

// NewDB returns a reference to the database connection.
func NewDB() *sql.DB {
	db, err := GetConnectionHandler()
	if err != nil {
		log.Fatalf("Failed to establish database connection: %v", err)
	}
	return db
}
