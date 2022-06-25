package authee

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"log"
	"suiibell/anatomy/authAnatomy"
	"suiibell/dbconn"
	"suiibell/ent"
	"suiibell/ent/user"
	"time"
)

func Register(registerUser ent.User) (authAnatomy.User, error) {
	var newUser authAnatomy.User

	db, errOpen := dbconn.DBConnection()
	if errOpen != nil {
		return newUser, errors.New("failed to open the database")
	}
	defer db.Close()

	ctx := context.Background()

	theUser := db.User.Query().Where(user.Email(registerUser.Email))
	isRegistered, errExist := theUser.Exist(ctx)
	if errExist != nil {
		log.Println("failed to query the user : ", errExist)
		return newUser, errors.New("failed to check the user")
	}

	if isRegistered {
		log.Println("the user is already registered")
		return newUser, errors.New("the user is already registered")
	}

	registerUser.ID = uuid.New()
	registerUser.IsBlocked = false
	registerUser.CreatedAt = time.Now().UTC()
	registerUser.UpdatedAt = time.Now().UTC()
	registerUser.FailedLoginAttempts = 0

	pass, errEncrypt := EncryptPass(string(registerUser.Password))
	if errEncrypt != nil {
		return newUser, errors.New("failed to encrypt the password")
	}

	_, errSave := db.User.Create().SetEmail(registerUser.Email).SetPassword(pass).Save(ctx)
	if errSave != nil {
		log.Println("failed to save the registerUser : ", errSave)
		return newUser, errSave
	}

	newUser.Email = registerUser.Email
	newUser.ID = registerUser.ID.String()
	newUser.CreatedAt = registerUser.CreatedAt.UTC()
	newUser.UpdatedAt = registerUser.UpdatedAt.UTC()
	newUser.FailedLoginAtttempts = registerUser.FailedLoginAttempts

	return newUser, nil
}
