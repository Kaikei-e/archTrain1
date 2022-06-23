package authee

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"suiibell/dbconn"
	"suiibell/ent"
	"suiibell/ent/user"
)

func LoginCheck(checkUser ent.User) (string, error) {

	db, errOpen := dbconn.DBConnection()
	if errOpen != nil {
		return "", errors.New("failed to open the database")
	}
	defer db.Close()

	ctx := context.Background()

	user, errQuery := db.User.Query().Where(user.Email(checkUser.Email)).Only(ctx)
	if errQuery != nil {
		return "", errQuery
	}

	isSuccess, errCompare := CompareHashedPassAndInput([]byte(user.Password), []byte(checkUser.Password))
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

func Register(user ent.User) error {

	db, errOpen := dbconn.DBConnection()
	if errOpen != nil {
		return errors.New("failed to open the database")
	}
	defer db.Close()

	ctx := context.Background()

	_, errSave := db.User.Create().SetID(uuid.New()).SetEmail(user.Email).SetPassword(user.Password).Save(ctx)
	if errSave != nil {
		return errSave
	}

	return nil
}
