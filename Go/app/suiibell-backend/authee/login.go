package authee

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"suiibell/dbconn"
	"suiibell/ent/user"
)

func LoginCheck(email, password string) (string, error) {

	db, errOpen := dbconn.DBConnection()
	if errOpen != nil {
		return "", errors.New("failed to open the database")
	}
	defer db.Close()

	ctx := context.Background()

	user, errQuery := db.User.Query().Where(user.Email(email)).Only(ctx)
	if errQuery != nil {
		return "", errQuery
	}

	isSuccess, errCompare := CompareHashedPassAndInput([]byte(user.Password), []byte(password))
	if errCompare != nil {
		return "", errors.New("failed to compare the hashed password and input password")
	}

	if !isSuccess {
		return "", errors.New("failed to compare the hashed password and input password")
	}

	if user.IsBlocked {
		return "", errors.New("the user is blocked")
	}

	return user.Email, nil
}

func Register(email, password string) error {
	//var user ent.User

	db, errOpen := dbconn.DBConnection()
	if errOpen != nil {
		return errors.New("failed to open the database")
	}
	defer db.Close()

	ctx := context.Background()

	_, errSave := db.User.Create().SetID(uuid.New()).SetEmail(email).SetPassword(password).Save(ctx)
	if errSave != nil {
		return errSave
	}

	return nil
}
