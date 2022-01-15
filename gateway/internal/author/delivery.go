package author

import "github.com/labstack/echo/v4"

type HttpDelivery interface {
	CreateAuthor() echo.HandlerFunc
}
