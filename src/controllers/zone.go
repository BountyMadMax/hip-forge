package controllers

import (
	"hip-forge/src/hetzner/dns_api"
	"hip-forge/src/models"
	"hip-forge/src/views"
	"hip-forge/src/views/pages"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Zones(c echo.Context) error {
	token := c.FormValue("token")

	if len(token) == 0 {
		return views.Render(c, http.StatusBadRequest, pages.ZoneOptions(nil))
	}

	api_zones, err := dns_api.GetZones(token)

	if err != nil {
		return views.Render(c, http.StatusInternalServerError, pages.ZoneOptions(nil))
	}

	zones := make([]models.Zone, len(api_zones))

	for i, zone := range api_zones {
		zones[i] = models.Zone{
			ID:   zone.Id,
			Name: zone.Name,
		}
	}

	return views.Render(c, http.StatusOK, pages.ZoneOptions(zones))
}
