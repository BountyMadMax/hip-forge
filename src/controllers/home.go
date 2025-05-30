package controllers

import (
	"hip-forge/src/database"
	"hip-forge/src/models"
	"hip-forge/src/views"
	"hip-forge/src/views/pages"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomeIndex(c echo.Context) error {
	accounts := []models.Account{}

	db := database.OpenDatabase()

	db.Find(&accounts)

	return views.Render(c, http.StatusOK, pages.Home(accounts, db))
}
