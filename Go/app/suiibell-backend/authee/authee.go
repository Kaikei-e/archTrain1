package authee

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"io/ioutil"
)

func LoginManager(e echo.Context) {
	userid := e.Bind("email")
	password := e.Bind("password")

	by, errRead := ioutil.ReadFile("./.env.rsa")

	jwt.ParseRSAPrivateKeyFromPEM()
}
