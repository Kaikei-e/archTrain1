package authee

import (
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"suiibell/anatomy/authAnatomy"
	"suiibell/ent"
	"time"
)

func LoginManager(e echo.Context) error {
	var incomeUserInfo authAnatomy.UserCredential
	var user ent.User

	errBind := e.Bind(&incomeUserInfo)
	if errBind != nil {
		return errors.New("failed to bind the user")
	}

	user.Email = incomeUserInfo.Email
	user.Password = incomeUserInfo.Password

	by, errRead := ioutil.ReadFile("./pkcs8.key")
	if errRead != nil {
		log.Println(errors.New("failed to read the rsa file."))
		return e.JSON(500, map[string]string{
			"error": "internal server error",
		})
	}

	authedId, errAuth := LoginCheck(user)
	if errAuth != nil {
		log.Println(errAuth)
		log.Printf("Incoming IP : %s", e.RealIP())
		return e.JSON(http.StatusBadRequest, map[string]string{
			"error": "failed to login the user",
		})
	}

	claims := jwt.MapClaims{
		"email": authedId,
		"exp":   time.Now().Add(time.Minute * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedString, err := token.SignedString(by)
	if err != nil {
		return errors.New("failed to sign the token")
	}

	return e.JSON(200, map[string]string{
		"token": signedString,
	})

}

func RegisterManager(e echo.Context) error {
	var incomeUserInfo authAnatomy.UserCredential
	var user ent.User

	errBind := e.Bind(&incomeUserInfo)
	if errBind != nil {
		return errors.New("failed to bind the user")
	}

	user.Email = incomeUserInfo.Email
	user.Password = incomeUserInfo.Password

	registeredUser, errRegister := Register(user)
	if errRegister != nil {
		return errRegister
	}

	by, errRead := ioutil.ReadFile("./pkcs8.key")
	if errRead != nil {
		return errors.New("failed to read the rsa file.")
	}

	signedKey, errKey := jwt.ParseRSAPrivateKeyFromPEM(by)
	if errKey != nil {
		return errors.New("failed to parse the rsa private key")
	}

	claims := jwt.MapClaims{
		"email": registeredUser.Email,
		"exp":   time.Now().Add(time.Minute * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedString, errSign := token.SignedString(signedKey)
	if errSign != nil {
		return errors.New("failed to sign the token")
	}

	return e.JSON(200, map[string]string{
		"token": signedString,
	})
}
