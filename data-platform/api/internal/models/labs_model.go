package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Lab struct {
	ID        	int
	Username  	string
	Location  	string
	CreatedAt 	*time.Time
	UpdatedAt 	*time.Time
	RanAt 		string
	Description string
}

// checkErr logs the error with additional context and terminates the program for critical errors.
func checkErr(err error, context string) {
	if err != nil {
		log.Fatalf("Context: %s, Error: %v", context, err)
	}
}

// createTable creates a new table in the database
func CreateLabsTable(db *sql.DB) error {
	createLabTableSQL := `
	CREATE TABLE IF NOT EXISTS labs (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) NOT NULL,
		location VARCHAR(50) NOT NULL,
		CreatedAt TIMESTAMPTZ DEFAULT NOW(),
		UpdatedAt TIMESTAMPTZ DEFAULT NOW(),
		RanAt  TIMESTAMPTZ DEFAULT NOW(),
		Description VARCHAR(50) NOT NULL
	);`

	_, err := db.Exec(createLabTableSQL)
	checkErr(err, "creating lab table")

	log.Println("User and Lab tables created successfully.")

	return nil
}

// GetAllLabs retrieves all labs from the labs table
func GetAllLabs(db *sql.DB) ([]Lab, error) {
	getLabsSQL := `
	SELECT id, username, location, CreatedAt, UpdatedAt, RanAt, Description
	FROM labs;`

	rows, err := db.Query(getLabsSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var labs []Lab
	for rows.Next() {
		var lab Lab
		err := rows.Scan(&lab.ID, &lab.Username, &lab.Location, &lab.CreatedAt, &lab.UpdatedAt, &lab.RanAt, &lab.Description)
		if err != nil {
			return nil, err
		}
		labs = append(labs, lab)
	}

	return labs, nil
}

func InsertLab(db *sql.DB, username, location string) error {
	escapedUsername := escapeString(username)
	escapedLocation := escapeString(location)

	// Create the SQL query with embedded values
	insertLabSQL := fmt.Sprintf(`
	INSERT INTO users (username, password)
	VALUES ('%s', '%s');`, escapedUsername, escapedLocation)

	_, err := db.Exec(insertLabSQL)
	checkErr(err, "inserting lab")

	log.Println("Lab inserted successfully.")

	return nil
}
