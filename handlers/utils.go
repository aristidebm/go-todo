package handlers

import (
    "context"
    "github.com/labstack/echo/v4"
	"github.com/a-h/templ"
    "test/todo/templates"
)

func Render(ctx echo.Context, tpl templ.Component) error{
    ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
    return tpl.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func renderServerError(ctx echo.Context, stackTrace string) error {
    component := templates.ServerError("500 Internal Error" , stackTrace)
    return component.Render(context.Background(), ctx.Response().Writer)
}

func renderValidationError(ctx echo.Context, message string) error {
    component := templates.ValidationError("400 Bad Request" , message)
    return component.Render(context.Background(), ctx.Response().Writer)
}

func renderNotFoundError(ctx echo.Context, message string) error {
    component := templates.ValidationError("404 Not found" , message)
    return component.Render(context.Background(), ctx.Response().Writer)
}
