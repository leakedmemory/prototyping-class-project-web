package db

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"

	"website/internal/models"
)

type DBI interface {
	Shutdown()

	AddUser(user *models.User) error
	GetUser(id string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id string) error

	AddPet(pet *models.Pet) error
	GetPet(id string) (*models.Pet, error)
	UpdatePet(pet models.Pet) error
	DeletePet(id string) error
}

type DB struct {
	userFile *os.File
	petFile  *os.File
	userData map[string]models.User
	petData  map[string]models.Pet
	mutex    sync.RWMutex
}

func NewDB(userFilePath, petFilePath string) (*DB, error) {
	userFile, err := os.OpenFile(userFilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	petFile, err := os.OpenFile(petFilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		userFile.Close()
		return nil, err
	}

	db := &DB{
		userFile: userFile,
		petFile:  petFile,
		userData: make(map[string]models.User),
		petData:  make(map[string]models.Pet),
	}

	err = db.load()
	if err != nil {
		db.Shutdown()
		return nil, err
	}

	return db, nil
}

func (db *DB) Shutdown() {
	db.userFile.Close()
	db.petFile.Close()
}

func (db *DB) load() error {
	userData, err := io.ReadAll(db.userFile)
	if err != nil {
		return err
	}

	if len(userData) > 0 {
		err = json.Unmarshal(userData, &db.userData)
		if err != nil {
			return err
		}
	}

	petData, err := io.ReadAll(db.petFile)
	if err != nil {
		return err
	}

	if len(petData) > 0 {
		err = json.Unmarshal(petData, &db.petData)
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) writeUserData() error {
	data, err := json.MarshalIndent(db.userData, "", "  ")
	if err != nil {
		return err
	}

	err = db.userFile.Truncate(0)
	if err != nil {
		return err
	}

	_, err = db.userFile.Seek(0, io.SeekStart)
	if err != nil {
		return nil
	}

	_, err = db.userFile.Write(data)
	return err
}

func (db *DB) AddUser(user *models.User) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if _, exists := db.userData[user.ID]; exists {
		return errors.New("user with this ID already exists")
	}

	db.userData[user.ID] = *user
	return db.writeUserData()
}

func (db *DB) GetUser(id string) (*models.User, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	user, ok := db.userData[id]
	if !ok {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (db *DB) UpdateUser(user *models.User) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if _, exists := db.userData[user.ID]; !exists {
		return errors.New("user not found")
	}

	db.userData[user.ID] = *user
	return db.writeUserData()
}

func (db *DB) DeleteUser(id string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if _, exists := db.userData[id]; !exists {
		return errors.New("user not found")
	}

	delete(db.userData, id)
	return db.writeUserData()
}

func (db *DB) writePetData() error {
	data, err := json.MarshalIndent(db.petData, "", "  ")
	if err != nil {
		return err
	}

	err = db.petFile.Truncate(0)
	if err != nil {
		return err
	}

	_, err = db.petFile.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	_, err = db.petFile.Write(data)
	return err
}

func (db *DB) AddPet(pet *models.Pet) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	user, ok := db.userData[pet.OwnerID]
	if !ok {
		return errors.New("user not found")
	}

	if _, exists := db.petData[pet.ID]; exists {
		return errors.New("pet with this ID already exists")
	}

	db.petData[pet.ID] = *pet
	user.PetCount++

	err := db.writePetData()
	if err != nil {
		return err
	}

	err = db.writeUserData()
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) GetPet(id string) (*models.Pet, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	pet, ok := db.petData[id]
	if !ok {
		return nil, errors.New("pet not found")
	}

	return &pet, nil
}

func (db *DB) UpdatePet(pet models.Pet) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if _, exists := db.petData[pet.ID]; !exists {
		return errors.New("pet not found")
	}

	db.petData[pet.ID] = pet
	return db.writePetData()
}

func (db *DB) DeletePet(id string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if _, exists := db.petData[id]; !exists {
		return errors.New("pet not found")
	}

	delete(db.petData, id)
	return db.writePetData()
}
