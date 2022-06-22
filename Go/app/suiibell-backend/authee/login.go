package authee

import (
	"errors"
	"suiibell/anatomy/authAnatomy"
	"suiibell/dbconn"
)

func LoginCheck(userid, password string) (string, error) {
	var user authAnatomy.User

	db, errOpen := dbconn.DBConnection()
	if errOpen != nil {
		return "", errors.New("failed to open the database")
	}

	tx := db.Where("email = ?", userid).First(&user)
	if tx.Error != nil {
		return "", errors.New("failed to find the user")
	}

	isSuccess, errCompare := CompareHashedPassAndInput([]byte(user.Password), []byte(password))
	if errCompare != nil {
		return "", errors.New("failed to compare the hashed password and input password")
	}

	if !isSuccess {
		return "", errors.New("failed to compare the hashed password and input password")
	}

	if user.FailedStatus {
		return "", errors.New("the user is blocked")
	}

	return user.Email, nil
}

func Register(userid, password string) error {
	var user authAnatomy.User

	db, errOpen := dbconn.DBConnection()
	if errOpen != nil {
		return errors.New("failed to open the database")
	}

	db.Where("email = ?", userid).First(&user)
	if user.Email != "" {
		return errors.New("the user already exists")
	}

	tx := db.Create(&user)
	if tx.Error != nil {
		return errors.New("failed to create the user")
	}

	return nil
}
