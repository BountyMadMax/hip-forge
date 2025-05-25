package controllers

import (
	"hip-forge/src/models"
	"hip-forge/src/views"
	"hip-forge/src/views/pages"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AccountNew(c echo.Context) error {
	account := models.Account{}

	return views.Render(c, http.StatusOK, pages.Account(account))
}

func AccountCreate(c echo.Context) error {
	account := models.Account{}

	return views.Render(c, http.StatusOK, pages.Account(account))
}

func AccountToggleTokenInput(c echo.Context) error {
	token := c.FormValue("token")
	hidden, err := strconv.ParseBool(c.FormValue("hidden"))
	if err != nil {
		hidden = true
	}

	return views.Render(c, http.StatusOK, pages.AccountTokenInput(token, !hidden))
}
