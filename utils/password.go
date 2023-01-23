package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes password using bcrypt at default cost i.e. 10
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}

// CheckPasswordHash compares hash with the supplied password
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

// CheckCommonPassword
