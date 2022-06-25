package authee

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// EncryptPass encrypts the password with bcrypt
func EncryptPass(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to encrypt the password")
	}

	return hashedPassword, nil

}

// CompareHashedPassAndInput compares the hashed password and input password
func CompareHashedPassAndInput(hashedPass, inputPass []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hashedPass, inputPass)
	if err != nil {
		return false, errors.New("failed to compare the hashed password and input password")
	}

	return true, nil

}
