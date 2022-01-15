package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *authorsHandlers) MapRoutes() {
	h.group.POST("", h.CreateAuthor())
	h.group.Any("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})
}
