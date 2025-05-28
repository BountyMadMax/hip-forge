package controllers

import (
	"hip-forge/src/hetzner/dns_api"
	"hip-forge/src/models"
	"hip-forge/src/views"
	"hip-forge/src/views/pages"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RecordNew(c echo.Context) error {
	record := models.Record{}

	return views.Render(c, http.StatusOK, pages.RecordRow(record))
}

func Records(c echo.Context) error {
	token := c.FormValue("token")
	zone := c.FormValue("zone")

	if len(zone) == 0 || len(token) == 0 {
		return views.Render(c, http.StatusBadRequest, pages.RecordRows(nil))
	}

	api_records, err := dns_api.GetRecords(token, &zone)

	if err != nil {
		return views.Render(c, http.StatusInternalServerError, pages.RecordRows(nil))
	}

	records := make([]models.Record, len(api_records))

	for i, record := range api_records {
		records[i] = models.Record{
			ID:    record.Id,
			Name:  record.Name,
			Value: record.Value,
		}
	}

	return views.Render(c, http.StatusOK, pages.RecordRows(records))
}
