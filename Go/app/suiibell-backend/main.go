package main

import (
	"log"
	"net/http"
	"suiibell/initRun"
	"suiibell/router"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	initRun.Init()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	errAuth := router.AuthRouting(e)
	log.Printf("failed to register the auth routing: %v", errAuth)

	//api := e.Group("/api")
	//api.Use(func(api echo.HandlerFunc) echo.HandlerFunc {
	//	return func(c echo.Context) error {
	//		err := router.LoginRouting(e)
	//		if err != nil {
	//			return err
	//		}
	//		return nil
	//	}
	//})

	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
