package controllers

import (
	"app/db"
	"app/models"
	"log"
)

//////////////////////////
// DB Methods
//////////////////////////

//
// Create

func CreateUser(user *models.User) error {
	return db.PgDB.Create(user).Error
}

//
// Read

// GetUserByEmail searches db for user with given email
// returns User, nil if found else nil, error
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	if err := db.PgDB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByUsername searches db for user with given username
// returns User, nil if found else nil, error
func GetUserByUsername(username string) (*models.User, error) {
	var user models.User

	if err := db.PgDB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

//
// Update

//
// Delete

//
// Checks

// Check if user exists or not (error is not handled)
func CheckUserExistsByEmail(email string) bool {

	if _, err := GetUserByEmail(email); err == nil {
		log.Print(err)
		return true
	}

	return false
}

// Check if user exists or not (error is not handled)
func CheckUserExistsByUsername(username string) bool {

	if _, err := GetUserByUsername(username); err == nil {
		log.Print(err)
		return true
	}

	return false
}
