package main

import (
	app "hip-forge/src"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	app.Router(e)

	e.Logger.Fatal(e.Start(":8080"))
}
