package handlers

import (
    "github.com/labstack/echo/v4"
	"github.com/a-h/templ"
)

func Render(c echo.Context, t templ.Component) error{
    c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
    return t.Render(c.Request().Context(), c.Response().Writer)
}
