package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(e echo.Context) error {
	if err := e.Render(http.StatusOK, "index", nil); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("templating: %v", err))
	}

	return nil
}
