package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres password=123456789 dbname=testdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")

	// Create table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username TEXT NOT NULL,
		password TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table created successfully!")

	// Insert data
	username := "testuser"
	password := "testpassword"
	_, err = db.Exec(`INSERT INTO users (username, password) VALUES ($1, $2)`, username, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data inserted successfully!")

	// Query and print the data
	rows, err := db.Query(`SELECT id, username, password FROM users`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Users table contents:")
	for rows.Next() {
		var id int
		var username, password string
		err := rows.Scan(&id, &username, &password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Username: %s, Password: %s\n", id, username, password)
	}

	// Check for errors from iterating over rows
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
