package db

import (
	"os"
	"testing"

	"website/internal/models"
)

func setupTestDB(t *testing.T) (*DB, func()) {
	t.Helper()

	// create temporary files for testing
	userFile, err := os.CreateTemp("", "users_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp user file: %v", err)
	}

	petFile, err := os.CreateTemp("", "pets_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp pet file: %v", err)
	}

	// create the database
	testDB, err := NewDB(userFile.Name(), petFile.Name())
	if err != nil {
		t.Fatalf("Failed to create test DB: %v", err)
	}

	// clean up function to close files and remove them after the test
	cleanup := func() {
		testDB.Shutdown()
		os.Remove(userFile.Name())
		os.Remove(petFile.Name())
	}

	return testDB, cleanup
}

func TestAddUser(t *testing.T) {
	testDB, cleanup := setupTestDB(t)
	defer cleanup()

	user := &models.User{ID: "1", Name: "John Doe"}
	err := testDB.AddUser(user)
	if err != nil {
		t.Fatalf("Failed to add user: %v", err)
	}

	retrievedUser, err := testDB.GetUser("1")
	if err != nil {
		t.Fatalf("Failed to get user: %v", err)
	}

	if *retrievedUser != *user {
		t.Errorf("Expected user %v, got %v", user, retrievedUser)
	}
}

func TestUpdateUser(t *testing.T) {
	testDB, cleanup := setupTestDB(t)
	defer cleanup()

	user := &models.User{ID: "1", Name: "John Doe"}
	err := testDB.AddUser(user)
	if err != nil {
		t.Fatalf("Failed to add user: %v", err)
	}

	updatedUser := &models.User{ID: "1", Name: "Jane Doe"}
	err = testDB.UpdateUser(updatedUser)
	if err != nil {
		t.Fatalf("Failed to update user: %v", err)
	}

	retrievedUser, err := testDB.GetUser("1")
	if err != nil {
		t.Fatalf("Failed to get user: %v", err)
	}

	if *retrievedUser != *updatedUser {
		t.Errorf("Expected updated user %v, got %v", updatedUser, retrievedUser)
	}
}

func TestDeleteUser(t *testing.T) {
	testDB, cleanup := setupTestDB(t)
	defer cleanup()

	user := &models.User{ID: "1", Name: "John Doe"}
	err := testDB.AddUser(user)
	if err != nil {
		t.Fatalf("Failed to add user: %v", err)
	}

	err = testDB.DeleteUser("1")
	if err != nil {
		t.Fatalf("Failed to delete user: %v", err)
	}

	_, err = testDB.GetUser("1")
	if err == nil {
		t.Fatalf("Expected error when getting deleted user, got nil")
	}
}

func TestAddPet(t *testing.T) {
	testDB, cleanup := setupTestDB(t)
	defer cleanup()

	user := &models.User{ID: "1", Name: "John Doe"}
	err := testDB.AddUser(user)
	if err != nil {
		t.Fatalf("Failed to add user: %v", err)
	}

	pet := &models.Pet{ID: "1", Name: "Rex", OwnerID: "1"}
	err = testDB.AddPet(pet)
	if err != nil {
		t.Fatalf("Failed to add pet: %v", err)
	}

	retrievedPet, err := testDB.GetPet("1")
	if err != nil {
		t.Fatalf("Failed to get pet: %v", err)
	}

	if *retrievedPet != *pet {
		t.Errorf("Expected pet %v, got %v", pet, retrievedPet)
	}
}

func TestDeletePet(t *testing.T) {
	testDB, cleanup := setupTestDB(t)
	defer cleanup()

	user := &models.User{ID: "1", Name: "John Doe"}
	err := testDB.AddUser(user)
	if err != nil {
		t.Fatalf("Failed to add user: %v", err)
	}

	pet := &models.Pet{ID: "1", Name: "Rex", OwnerID: "1"}
	err = testDB.AddPet(pet)
	if err != nil {
		t.Fatalf("Failed to add pet: %v", err)
	}

	err = testDB.DeletePet("1")
	if err != nil {
		t.Fatalf("Failed to delete pet: %v", err)
	}

	_, err = testDB.GetPet("1")
	if err == nil {
		t.Fatalf("Expected error when getting deleted pet, got nil")
	}
}
