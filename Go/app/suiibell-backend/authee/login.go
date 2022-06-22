package authee

import (
	"errors"
	"suiibell/anatomy/authAnatomy"
	"suiibell/dbconn"
)

func LoginCheck(userid, password string) error {
	var user authAnatomy.User

	db, errOpen := dbconn.DBConnection()
	if errOpen != nil {
		return errors.New("failed to open the database")
	}

	tx := db.Where("email = ?", userid).First(&user)
	if tx.Error != nil {
		return errors.New("failed to find the user")
	}

	isSuccess, errCompare := CompareHashedPassAndInput([]byte(user.Password), []byte(password))
	if errCompare != nil {
		return errors.New("failed to compare the hashed password and input password")
	}

	return nil
}
