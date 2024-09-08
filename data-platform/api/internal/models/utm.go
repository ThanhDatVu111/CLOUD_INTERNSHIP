package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type UtmDevice struct {
	ID                 int
	CreatedTime        time.Time
	MaxTensileStrength float64
}

func NewUtmDevice(id int, time time.Time, maxTensileStrength float64) UtmDevice {
	return UtmDevice{
		ID:                 id,
		CreatedTime:        time,
		MaxTensileStrength: maxTensileStrength,
	}
}

// createTable creates a new table in the database
func createTable(db *sql.DB, tableName string) error {
	createTableSQL := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		ID SERIAL PRIMARY KEY,
		CreatedTime TIMESTAMP NOT NULL,
		MaxTensileStrength FLOAT NOT NULL
	);`, tableName)

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Printf("Error creating table: %v", err)
		return fmt.Errorf("error creating table: %v", err)
	}

	log.Println("Table created successfully.")
	return nil
}

func InsertUtmDevice(db *sql.DB, device UtmDevice) error {
	/*
		Insert the device into a devices table. SQL insert data
		Postgres
	*/
	query := fmt.Sprintf(`
	INSERT INTO devices (ID, CreatedTime, MaxTensileStrength)
	VALUES (%d, '%s', %.10f)
	RETURNING ID;`, device.ID, device.CreatedTime.Format(time.RFC3339), device.MaxTensileStrength)

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error inserting device: %v", err)
	}

	log.Printf("Device inserted successfully with ID: %d", device.ID)
	return nil
}

func GetUtmDevice(db *sql.DB, id int) (UtmDevice, error) {
	// Create the SQL query string using fmt.Sprintf
	selectSQL := fmt.Sprintf(`SELECT ID, CreatedTime, MaxTensileStrength FROM devices WHERE ID = %d;`, id)

	var device UtmDevice
	// Execute the query without passing the `id` parameter separately
	err := db.QueryRow(selectSQL).Scan(&device.ID, &device.CreatedTime, &device.MaxTensileStrength)
	if err != nil {
		if err == sql.ErrNoRows {
			return device, fmt.Errorf("device with ID %d not found", id)
		}
		log.Printf("Error retrieving device: %v", err)
		return device, fmt.Errorf("error retrieving device: %v", err)
	}

	log.Printf("Device retrieved successfully: %+v", device)
	return device, nil
}

func DeleteUtmDevice(db *sql.DB, id int) error {
	/*
		Delete the device from the devices table. SQL delete data
		Postgres
	*/
	deleteSQL := fmt.Sprintf(`DELETE FROM devices WHERE ID = %d;`, id)

	result, err := db.Exec(deleteSQL, id)
	if err != nil {
		log.Printf("Error deleting device: %v", err)
		return fmt.Errorf("error deleting device: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("device with ID %d not found", id)
	}

	log.Printf("Device deleted successfully with ID: %d", id)
	return nil
}

func UpdateUtmDevice(db *sql.DB, device UtmDevice) error {
	/*
		Update the device in the devices table. SQL update data
		Postgres
	*/
	updateSQL := fmt.Sprintf(`
	UPDATE devices
	SET MaxTensileStrength = %f
	WHERE ID = %d;`, device.MaxTensileStrength, device.ID)

	result, err := db.Exec(updateSQL, device.CreatedTime, device.MaxTensileStrength, device.ID)
	if err != nil {
		log.Printf("Error updating device: %v", err)
		return fmt.Errorf("error updating device: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("device with ID %d not found", device.ID)
	}

	log.Printf("Device updated successfully with ID: %d", device.ID)
	return nil
}

/*
class NewUtmDevice:
	def __init__(self, ID, Time, Strength):
		self.ID = ID
		self.Time = Time
		self.MaxTensileStrength = Strength

device1 = NewUtmDevice(ID = 1, Time = difftime, Strength = 20.0)
device2 = NewUtmDevice(ID = 2, Time = newtime, Strength = 10.0)

*/
