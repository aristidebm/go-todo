package handlers

import (
	"context"
	"log"

	"github.com/labstack/echo/v4"

	"test/todo/database"
	"test/todo/models"
	"test/todo/templates"
)

func ListTodo(ctx echo.Context) error {

    db, err := database.OpenDatabase()

    if err != nil {
        return renderServerError(ctx, err.Error()) 
    }

    log.Println("listing todos ...") 
    
    store := database.NewTodoStore(db)

    var todos []models.Todo 

    if todos, err = store.ListTodos(); err != nil {
        return renderServerError(ctx, err.Error()) 
    }

    component := templates.TodoList("List of todos" , todos)

    return component.Render(context.Background(), ctx.Response().Writer)
}
