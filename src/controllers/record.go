package controllers

import (
	"hip-forge/src/models"
	"hip-forge/src/views"
	"hip-forge/src/views/pages"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RecordNew(c echo.Context) error {
	record := models.DNSRecord{}

	return views.Render(c, http.StatusOK, pages.DnsRecordRow(record))
}
