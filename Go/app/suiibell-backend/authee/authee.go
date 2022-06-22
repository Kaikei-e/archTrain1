package authee

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"io/ioutil"
	"suiibell/dbconn"
)

func LoginManager(e echo.Context) error {
	userid := e.Bind("email")
	password := e.Bind("password")

	by, errRead := ioutil.ReadFile("./id_rsa")
	if errRead != nil {
		return errors.New("failed to read the rsa file.")
	}

	db, errOpen := dbconn.DBConnection()
	if errOpen != nil {

	}

	rsa_pk, errParse := jwt.ParseRSAPrivateKeyFromPEM(by)
	if errParse != nil {
		return errors.New("failed to parse the rsa private key")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"userid":   userid,
		"password": password,
	})

	return nil

}
