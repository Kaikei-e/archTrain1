package authee

import (
	"context"
	"errors"
	"log"
	"suiibell/dbconn"
	"suiibell/ent"
	"suiibell/ent/user"
)

func LoginCheck(checkUser ent.User) (string, error) {

	log.Println("checkUser : ", checkUser)

	db, errOpen := dbconn.DBConnection()
	if errOpen != nil {
		return "", errors.New("failed to open the database")
	}
	defer db.Close()

	ctx := context.Background()

	theUser, errQuery := db.User.Query().Where(user.Email(checkUser.Email)).Only(ctx)
	if errQuery != nil {
		return "", errors.New("failed to query the user")
	}

	isSuccess, errCompare := CompareHashedPassAndInput([]byte(theUser.Password), []byte(checkUser.Password))
	if errCompare != nil {
		return "", errors.New("failed to compare the hashed password and input password")
	}

	if !isSuccess {
		return "", errors.New("failed to compare the hashed password and input password")
	}

	if theUser.IsBlocked {
		return "", errors.New("the only is blocked")
	}

	return theUser.Email, nil
}
