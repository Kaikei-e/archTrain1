package authee

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"io/ioutil"
	"suiibell/ent"
	"time"
)

func LoginManager(e echo.Context) error {
	var user ent.User

	errBind := e.Bind(&user)
	if errBind != nil {
		return errors.New("failed to bind the user")
	}

	by, errRead := ioutil.ReadFile("./id_rsa")
	if errRead != nil {
		return errors.New("failed to read the rsa file.")
	}

	authedId, errAuth := LoginCheck(user)
	if errAuth != nil {
		return errAuth
	}

	claims := jwt.MapClaims{
		"email": authedId,
		"exp":   time.Now().Add(time.Minute * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	//
	//rsa_pk, errParse := jwt.ParseRSAPrivateKeyFromPEM(by)
	//if errParse != nil {
	//	return errors.New("failed to parse the rsa private key")
	//}

	signedString, err := token.SignedString(by)
	if err != nil {
		return errors.New("failed to sign the token")
	}

	return e.JSON(200, map[string]string{
		"token": signedString,
	})

}

func RegisterManager(e echo.Context) error {
	var user ent.User

	errBind := e.Bind(&user)
	if errBind != nil {
		return errors.New("failed to bind the user")
	}

	fmt.Println(user)

	by, errRead := ioutil.ReadFile("./suiibell_rsa")

	if errRead != nil {
		return errors.New("failed to read the rsa file.")
	}

	registeredUser, errRegister := Register(user)

	if errRegister != nil {
		return errors.New("failed to register the user")
	}

	claims := jwt.MapClaims{
		"email": registeredUser.Email,
		"exp":   time.Now().Add(time.Minute * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedKey, errKey := jwt.ParseRSAPrivateKeyFromPEM(by)
	if errKey != nil {
		return errors.New("failed to parse the rsa private key")
	}

	signedString, errSign := token.SignedString(signedKey)
	if errSign != nil {
		return errors.New("failed to sign the token")
	}

	return e.JSON(200, map[string]string{
		"token": signedString,
	})
}
