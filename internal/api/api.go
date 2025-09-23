package api

import (
	"net/http"

	"github.com/gonzalocarreteroh/weather-api/internal/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Weather *service.WeatherService
}

func (h *Handler) GetWeather(c echo.Context) error {
	city := c.Param("city")
	wr, err := h.Weather.GetWeather(c.Request().Context(), city)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, wr)
}
