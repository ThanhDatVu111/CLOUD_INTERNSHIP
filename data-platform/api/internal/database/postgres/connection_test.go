package postgres

import (
	"api/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConnectionHandler(t *testing.T) {
	db, err := GetConnectionHandler()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		t.Fatalf("Failed to ping the database: %v", err)
	}
	t.Log("Successfully connected and pinged the database!")
}

func TestDatabaseOperations(t *testing.T) {
	db, err := GetConnectionHandler()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		t.Fatalf("Failed to ping the database: %v", err)
	}

	err = models.InsertUser(db, "john", "123")
	assert.NoError(t, err, "There was an issue with inserting the user")

	err = models.InsertLab(db, "sarah", "url")
	assert.NoError(t, err, "There was an issue with inserting the lab")

	users, err := models.GetAllUsers(db, 100)
	assert.NoError(t, err, "There was an issue with getting all the users")
	assert.NotEqual(t, len(users), 0, "Len of users greater than 0")

	labs, err := models.GetAllLabs(db)
	assert.NoError(t, err, "There was an issue with getting all the labs")
	assert.NotEqual(t, len(labs), 0, "Len of labs greater than 0")

}
