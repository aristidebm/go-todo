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
	router.POST("/todos", handlers.CreateTodo)
	router.GET("/todos/:id", handlers.RetrieveTodo)

	log.Println("handlers are setup")
	// router.PATCH("/todos/:id", handlers.UpdateTodo)
	// router.DELETE("/todos/:id", handlers.DeleteTodo)

	if err := database.Migrate(); err != nil {
		log.Fatalf("cannot initialize the database: %v", err)
	}

	log.Println("database is setup")

	router.Logger.Fatal(router.Start(":8000"))
}
