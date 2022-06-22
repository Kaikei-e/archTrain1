package authee

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"io/ioutil"
	"time"
)

func LoginManager(e echo.Context) error {
	userid := e.FormValue("email")
	password := e.FormValue("password")

	by, errRead := ioutil.ReadFile("./id_rsa")
	if errRead != nil {
		return errors.New("failed to read the rsa file.")
	}

	authedId, errAuth := LoginCheck(userid, password)
	if errAuth != nil {
		return errAuth
	}

	claims := jwt.MapClaims{
		"userid": authedId,
		"exp":    time.Now().Add(time.Minute * 30).Unix(),
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
	userid := e.FormValue("email")
	password := e.FormValue("password")

	by, errRead := ioutil.ReadFile("./id_rsa")

	if errRead != nil {
		return errors.New("failed to read the rsa file.")
	}

	errRegister := Register(userid, password)

	if errRegister != nil {
		return errors.New("failed to register the user")
	}

	claims := jwt.MapClaims{
		"userid": userid,
		"exp":    time.Now().Add(time.Minute * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedString, errSign := token.SignedString(by)
	if errSign != nil {
		return errors.New("failed to sign the token")
	}

	return e.JSON(200, map[string]string{
		"token": signedString,
	})
}
