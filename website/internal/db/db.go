package db

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"

	"website/internal/models"
)

type DB struct {
	userFile *os.File
	userData map[string]models.User
	mutex    sync.RWMutex
}

func NewDB(userFilePath string) (*DB, error) {
	userFile, err := os.OpenFile(userFilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	db := &DB{
		userFile: userFile,
		userData: make(map[string]models.User),
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

func (db *DB) GetUserByID(id string) (*models.User, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	user, exists := db.userData[id]
	if exists {
		return &user, nil
	} else {
		return nil, errors.New("user not found")
	}
}

func (db *DB) GetUserByEmailAndPassword(email, password string) (*models.User, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	for _, user := range db.userData {
		if user.Email == email {
			if user.Password == password {
				return &user, nil
			}
			return nil, errors.New("incorrect password")
		}
	}

	return nil, errors.New("user not found")
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
