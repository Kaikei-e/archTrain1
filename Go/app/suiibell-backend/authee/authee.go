package authee

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"io/ioutil"
)

func LoginManager(e echo.Context) error {
	userid := e.Bind("email")
	password := e.Bind("password")

	by, errRead := ioutil.ReadFile("./id_rsa")
	if errRead != nil {
		return errors.New("failed to read the rsa file.")
	}

	rsa_pk, errParse := jwt.ParseRSAPrivateKeyFromPEM(by)
	if errParse != nil {
		return errors.New("failed to parse the rsa private key")
	}

}
