package authee

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"log"
	"suiibell/anatomy/authAnatomy"
	"suiibell/dbconn"
	"suiibell/ent"
	"suiibell/ent/user"
	"time"
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

func Register(user ent.User) (authAnatomy.User, error) {
	var newUser authAnatomy.User

	db, errOpen := dbconn.DBConnection()
	if errOpen != nil {
		return newUser, errors.New("failed to open the database")
	}
	defer db.Close()

	ctx := context.Background()

	user.ID = uuid.New()
	user.IsBlocked = false
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()
	user.FailedLoginAttempts = 0

	pass, errEncrypt := EncryptPass(string(user.Password))
	if errEncrypt != nil {
		return newUser, errors.New("failed to encrypt the password")
	}

	_, errSave := db.User.Create().SetEmail(user.Email).SetPassword(pass).Save(ctx)
	if errSave != nil {
		log.Println("failed to save the user : ", errSave)
		return newUser, errSave
	}

	newUser.Email = user.Email
	newUser.ID = user.ID.String()
	newUser.CreatedAt = user.CreatedAt.UTC()
	newUser.UpdatedAt = user.UpdatedAt.UTC()
	newUser.FailedLoginAtttempts = user.FailedLoginAttempts

	return newUser, nil
}
