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

// LoginManager is the function to authenticate the user
func LoginManager(e echo.Context) error {
	var incomeUserInfo authAnatomy.UserCredential
	var user ent.User

	errBind := e.Bind(&incomeUserInfo)
	if errBind != nil {
		return errors.New("failed to bind the user")
	}

	user.Email = incomeUserInfo.Email
	user.Password = incomeUserInfo.Password

	// call the login function to authenticate the user and get the user email
	authedId, errAuth := LoginCheck(user)
	if errAuth != nil {
		log.Println(errAuth)
		log.Printf("Incoming IP : %s", e.RealIP())
		return e.JSON(http.StatusBadRequest, map[string]string{
			"error": "failed to login the user",
		})
	}

	by, errRead := ioutil.ReadFile("./pkcs8.key")
	if errRead != nil {
		log.Println(errors.New("failed to read the rsa file."))
		return e.JSON(500, map[string]string{
			"error": "internal server error",
		})
	}

	signedKey, errKey := jwt.ParseRSAPrivateKeyFromPEM(by)
	if errKey != nil {
		return errors.New("failed to parse the rsa private key")
	}

	claims := jwt.MapClaims{
		"email": authedId,
		"exp":   time.Now().Add(time.Minute * 30).Unix(),
	}

	jwtoken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedString, errSign := jwtoken.SignedString(signedKey)
	if errSign != nil {
		return errors.New("failed to sign the jwtoken")
	}

	return e.JSON(200, map[string]string{
		"email":   authedId,
		"jwtoken": signedString,
	})

}

// RegisterManager is the function to register the user
func RegisterManager(e echo.Context) error {

	// declare the user struct to store the user information for minimum information
	var incomeUserInfo authAnatomy.UserCredential
	var user ent.User

	errBind := e.Bind(&incomeUserInfo)
	if errBind != nil {
		return errors.New("failed to bind the user")
	}

	// pass the user information to the ent.user struct
	user.Email = incomeUserInfo.Email
	user.Password = incomeUserInfo.Password

	// call the register function to register the user
	registeredUser, errRegister := Register(user)
	if errRegister != nil {
		return errRegister
	}

	// file format must be the format of pkcs
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
