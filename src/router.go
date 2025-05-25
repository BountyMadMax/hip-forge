package router

import (
	"context"
	"hip-forge/src/views"
	"hip-forge/src/views/pages"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {
	addStaticFiles(e)

	e.GET("/", func(c echo.Context) error {
		return Render(c, http.StatusOK, pages.Home())
	})
}

func addStaticFiles(e *echo.Echo) {
	e.File("assets/js/htmx.min.js", "node_modules/htmx.org/dist/htmx.min.js")
	e.File("assets/css/main.css", "src/assets/css/output.css")
	e.File("assets/css/inter.css", "node_modules/@fontsource-variable/inter/index.css")
	e.Static("assets/font/inter/", "node_modules/@fontsource-variable/inter/files/")
	e.Static("assets/icons/", "node_modules/lucide-static/font/")
}

func Render(ctx echo.Context, HTTPStatus int, t templ.Component) error {
	// See https://htmx.org/docs/#caching
	ctx.Response().Writer.Header().Add("Vary", "HX-Request")

	ctx.Response().Writer.WriteHeader(HTTPStatus)

	var err error
	// Only return the fragment (the given Component) on a HTMX request, else wrap it in the layout.
	if ctx.Request().Header.Get("HX-Request") == "true" {
		err = t.Render(context.Background(), ctx.Response().Writer)
	} else {
		wrapped := templ.WithChildren(context.Background(), t)
		err = views.Layout().Render(wrapped, ctx.Response().Writer)
	}

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to render response")
	}

	return nil
}
