package models

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestNewUtmDevice(t *testing.T) {
	device1 := NewUtmDevice(1, time.Now(), 10.0)
	assert.Equal(t, device1.ID, 1, "The two values should be the same.")
}

const (
	dbDriver   = "postgres"
	dbSource   = "user=sanjithdevineni password=password dbname=postgres sslmode=disable host=127.0.0.1"
	testDBName = "platform_db"
	tableName  = "devices"
)

var testDB *sql.DB

func connectToDB(dataSourceName string) (*sql.DB, error) {
	var db *sql.DB
	var err error
	for i := 0; i < 5; i++ { // Retry up to 5 times
		db, err = sql.Open(dbDriver, dataSourceName)
		if err == nil {
			err = db.Ping()
			if err == nil {
				return db, nil
			}
		}
		time.Sleep(2 * time.Second) // This waits before retrying
	}
	return nil, err
}

func setupTable() {
	var err error
	testDB, err = connectToDB(dbSource)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create the test database
	_, err = testDB.Exec("CREATE DATABASE " + testDBName)
	if err != nil && err.Error() != "pq: database \""+testDBName+"\" already exists" {
		log.Fatalf("Failed to create test database: %v", err)
	}

	// Connect to the test database
	testDB, err = connectToDB(dbSource + " dbname=" + testDBName)
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}

	// Create the table in the test database
	err = createTable(testDB, tableName)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}

func teardownTable() {
	// Close the connection to the test database
	testDB.Close()

	// Reconnect to the main database to drop the test database
	mainDB, err := connectToDB(dbSource)
	if err != nil {
		log.Fatalf("Failed to reconnect to main database: %v", err)
	}
	defer mainDB.Close()

	_, err = mainDB.Exec("DROP DATABASE " + testDBName)
	if err != nil {
		log.Fatalf("Failed to drop test database: %v", err)
	}
	log.Println("Test database dropped successfully.")
}

func TestMain(m *testing.M) {
	setupTable()
	code := m.Run()
	teardownTable()
	os.Exit(code)
}

func TestInsertUtmDevice(t *testing.T) {
	device := NewUtmDevice(1, time.Now(), 10.0)
	err := InsertUtmDevice(testDB, device)
	assert.NoError(t, err, "Expected no error, got error while inserting device")
}

func TestGetUtmDevice(t *testing.T) {
	device := NewUtmDevice(2, time.Now(), 15.0)
	err := InsertUtmDevice(testDB, device)
	assert.NoError(t, err, "Expected no error, got error while inserting device")

	retrievedDevice, err := GetUtmDevice(testDB, device.ID)
	assert.NoError(t, err, "Expected no error, got error while retrieving device")
	assert.Equal(t, device.ID, retrievedDevice.ID, "The IDs should be the same")
	assert.Equal(t, device.MaxTensileStrength, retrievedDevice.MaxTensileStrength, "The MaxTensileStrength should be the same")
}

func TestDeleteUtmDevice(t *testing.T) {
	t.Skip("skipping delete test for now.")
	device := NewUtmDevice(3, time.Now(), 20.0)
	err := InsertUtmDevice(testDB, device)
	assert.NoError(t, err, "Expected no error, got error while inserting device")

	err = DeleteUtmDevice(testDB, device.ID)
	assert.NoError(t, err, "Expected no error, got error while deleting device")

	_, err = GetUtmDevice(testDB, device.ID)
	assert.Error(t, err, "Expected error, got no error while retrieving deleted device")
}

func TestUpdateUtmDevice(t *testing.T) {
	t.Skip("skipping update test for now.")
	device := NewUtmDevice(4, time.Now(), 25.0)
	err := InsertUtmDevice(testDB, device)
	assert.NoError(t, err, "Expected no error, got error while inserting device")

	updatedDevice := UtmDevice{
		ID:                 device.ID,
		CreatedTime:        time.Now(),
		MaxTensileStrength: 30.0,
	}
	err = UpdateUtmDevice(testDB, updatedDevice)
	assert.NoError(t, err, "Expected no error, got error while updating device")

	retrievedDevice, err := GetUtmDevice(testDB, updatedDevice.ID)
	assert.NoError(t, err, "Expected no error, got error while retrieving updated device")
	assert.Equal(t, updatedDevice.MaxTensileStrength, retrievedDevice.MaxTensileStrength, "The MaxTensileStrength should be updated")
}