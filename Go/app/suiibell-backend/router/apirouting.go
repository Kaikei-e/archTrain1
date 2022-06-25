// Package router for api routing
package router

import (
	"github.com/labstack/echo"
	"log"
	"suiibell/authee"
)

// AuthRouting is the function to handle the api routing
func AuthRouting(e *echo.Echo) error {
	e.POST("/api/v1/login", func(context echo.Context) error {
		err := authee.LoginManager(context)
		if err != nil {
			return err
		}
		return nil
	})

	e.POST("/api/v1/register", func(context echo.Context) error {
		err := authee.RegisterManager(context)
		if err != nil {
			log.Println("failed to register the user : ", err)
			return err
		}
		return nil
	})

	return nil
}
