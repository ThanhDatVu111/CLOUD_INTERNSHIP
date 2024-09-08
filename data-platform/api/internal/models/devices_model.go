package models

import (
	"database/sql"
	"fmt"
	"log"
)

// User struct represents a user record in the database
type Device struct {
	ID     int
    Name   string
}

// createTable creates a new table in the database
func CreateDevicesTable(db *sql.DB) error {
	createDeviceTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50) NOT NULL,
	);`

	_, err := db.Exec(createDeviceTableSQL)
    if err != nil {
        return fmt.Errorf("creating device table: %v", err)
    }

    log.Println("Device table created successfully.")
    return nil
}

// InsertDevice inserts a new device into the database
func InsertDevice(db *sql.DB, name string) error {
    // Escape the user inputs to prevent SQL injection
    escapedName := escapeString(name)

    // Create the SQL query with embedded values
    insertDeviceSQL := fmt.Sprintf(`
    INSERT INTO devices (name) VALUES ('%s');`, escapedName)

    _, err := db.Exec(insertDeviceSQL)
    if err != nil {
        return fmt.Errorf("inserting device: %v", err)
    }

    log.Println("Device inserted successfully.")
    return nil
}


// GetAllDevices retrieves the list of devices from the database
func GetAllDevices(db *sql.DB, limit int) ([]Device, error) {
    getDevicesSQL := fmt.Sprintf(`
	SELECT id, name
	FROM devices
	LIMIT %d;`, limit)

	rows, err := db.Query(getDevicesSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var devices []Device
	for rows.Next() {
		var device Device
		err := rows.Scan(&device.ID, &device.Name)
		if err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}
	return devices, nil
}
