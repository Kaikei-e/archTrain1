package main

import (
	"suiibell/initRun"
	"suiibell/router"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	err := initRun.Init()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	errAuth := router.AuthRouting(e)
	if errAuth != nil {
		panic(errAuth)
	}

	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}
