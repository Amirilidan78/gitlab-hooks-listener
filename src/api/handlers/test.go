package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleTest(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]any{
		"success": true,
	})
}