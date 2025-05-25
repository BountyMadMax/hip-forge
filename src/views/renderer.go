package views

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

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
		err = Layout().Render(wrapped, ctx.Response().Writer)
	}

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to render response")
	}

	return nil
}
