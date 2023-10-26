package main

import (
	"html/template"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

type Todo struct {
	Title    string
	Content  string
	Priority string
}

func main() {
	router := echo.New()
	todos := []Todo{
		{Title: "Touch Typing", Content: "My Learning touch typing with calmack-dh keybord", Priority: "High"},
		{Title: "English", Content: "I need to learn english", Priority: "Medium"},
		{Title: "French", Content: "I fluently speak french", Priority: "Low"},
	}
	router.GET("/todos", func(c echo.Context) error {
		// source https://www.reddit.com/r/golang/comments/17d12wk/using_echo_with_ahtempl/
		handler := templ.Handler(TodoList(todos))
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
		return handler.Component.Render(c.Request().Context(), c.Response().Writer)
	})
	router.Logger.Fatal(router.Start(":8000"))
}
