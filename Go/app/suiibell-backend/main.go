package main

import (
	"net/http"
	"suiibell/initRun"

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

	api := e.Group("/api")
	api.Use()

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
