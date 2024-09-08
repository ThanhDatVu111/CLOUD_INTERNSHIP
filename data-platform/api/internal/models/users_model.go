package models

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

// User struct represents a user record in the database
type User struct {
	ID       int
	Username string
	Password string
}

// createTable creates a new table in the database
func CreateUsersTable(db *sql.DB) error {
	createUserTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) NOT NULL,
		password VARCHAR(50) NOT NULL
	);`

	_, err := db.Exec(createUserTableSQL)
	checkErr(err, "creating user table")

	log.Println("User and Lab tables created successfully.")

	return nil
}

// Function to escape user inputs
func escapeString(s string) string {
	// Implement your escaping logic here (basic example)
	return strings.ReplaceAll(s, "'", "''")
}

func InsertUser(db *sql.DB, username, password string) error {
	// Escape the user inputs to prevent SQL injection
	escapedUsername := escapeString(username)
	escapedPassword := escapeString(password)

	// Create the SQL query with embedded values
	insertUserSQL := fmt.Sprintf(`
	INSERT INTO users (username, password)
	VALUES ('%s', '%s');`, escapedUsername, escapedPassword)

	_, err := db.Exec(insertUserSQL)
	checkErr(err, "inserting user")

	log.Println("User inserted successfully.")

	return nil
}


func GetAllUsers(db *sql.DB, limit int) ([]User, error) {
	getUsersSQL := fmt.Sprintf(`
	SELECT id, username, password
	FROM users
	LIMIT %d;`, limit)

	rows, err := db.Query(getUsersSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
