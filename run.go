package main

import (
	"log"
	"test/todo/database"
	"test/todo/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	router := echo.New()

	router.GET("/todos", handlers.ListTodo)
	router.GET("/todos", handlers.CreateTodo)
	router.GET("/todos/:id", handlers.RetrieveTodo)
	// router.GET("/todos/:id", handlers.UpdateTodo)
	// router.GET("/todos/:id", handlers.DeleteTodo)

	if err := database.Migrate(); err != nil {
		log.Fatalf("cannot initialize the database %v", err)
	}

	router.Logger.Fatal(router.Start(":8000"))
}
