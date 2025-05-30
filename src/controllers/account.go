package controllers

import (
	"hip-forge/src/database"
	"hip-forge/src/models"
	"hip-forge/src/views"
	"hip-forge/src/views/pages"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AccountNew(c echo.Context) error {
	account := models.Account{}

	return views.Render(c, http.StatusOK, pages.Account(account, database.OpenDatabase()))
}

func AccountSave(c echo.Context) error {
	name := c.FormValue("name")
	token := c.FormValue("token")
	zone := c.FormValue("zone")

	if len(name) < 1 || len(token) < 1 || len(zone) < 1 {
		return views.Render(c, http.StatusBadRequest, nil)
	}

	// Use MultipartForm because of the list of domains.
	multiForm, err := c.MultipartForm()

	if err != nil {
		return views.Render(c, http.StatusBadRequest, nil)
	}

	recordsId := multiForm.Value["id"]
	recordsDomain := multiForm.Value["domain"]
	recordsHidden := multiForm.Value["hidden"]

	if len(recordsId) != len(recordsDomain) && len(recordsDomain) != len(recordsHidden) {
		return views.Render(c, http.StatusBadRequest, nil)
	}

	records := make([]models.Record, len(recordsId))

	for i, id := range recordsId {
		hidden, err := strconv.ParseBool(recordsHidden[i])
		if err != nil {
			hidden = false
		}

		records[i] = models.Record{ID: id, Name: recordsDomain[i], Hidden: hidden}
	}

	db := database.OpenDatabase()

	// Save account.

	// Save zone.

	// Save records.

	accounts := []models.Account{}
	db.Find(&accounts)

	return views.Render(c, http.StatusOK, pages.Home(accounts, db))
}

func AccountToggleTokenInput(c echo.Context) error {
	token := c.FormValue("token")
	hidden, err := strconv.ParseBool(c.FormValue("hidden"))
	if err != nil {
		hidden = true
	}

	return views.Render(c, http.StatusOK, pages.AccountTokenInput(token, !hidden))
}

func AccountTokenChanged(c echo.Context) error {
	token := c.FormValue("token")

	return views.Render(c, http.StatusOK, pages.ZoneInput(len(token) > 0))
}
