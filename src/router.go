package app

import (
	"hip-forge/src/controllers"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {
	addStaticFiles(e)

	e.GET("/", func(c echo.Context) error {
		return controllers.HomeIndex(c)
	})

	e.POST("/accounts", func(c echo.Context) error {
		return controllers.AccountNew(c)
	})
	e.POST("/accounts/new", func(c echo.Context) error {
		return controllers.AccountCreate(c)
	})

	e.POST("/records", func(c echo.Context) error {
		return controllers.RecordNew(c)
	})

	e.POST("/accounts/toggle-token-input", func(c echo.Context) error {
		return controllers.AccountToggleTokenInput(c)
	})
}

func addStaticFiles(e *echo.Echo) {
	e.File("assets/js/htmx.min.js", "node_modules/htmx.org/dist/htmx.min.js")
	e.File("assets/css/main.css", "src/assets/css/output.css")
	e.File("assets/css/inter.css", "node_modules/@fontsource-variable/inter/index.css")
	e.Static("assets/font/inter/", "node_modules/@fontsource-variable/inter/files/")
	e.Static("assets/icons/", "node_modules/lucide-static/font/")
}
