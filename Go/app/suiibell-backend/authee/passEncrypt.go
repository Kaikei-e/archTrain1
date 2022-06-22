package authee

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPass(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to encrypt the password")
	}

	return hashedPassword, nil

}

func CompareHashedPassAndInput(hashedPass, inputPass []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hashedPass, []byte(inputPass))
	if err != nil {
		return false, errors.New("failed to compare the hashed password and input password")
	}

	return true, nil

}