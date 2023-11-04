package main

import (
	"github.com/labstack/echo/v4"
	"test/todo/handlers"
)

func main() {
	router := echo.New()
	router.GET("/todos", handlers.ListTodo)
	router.Logger.Fatal(router.Start(":8000"))
}
