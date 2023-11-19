package handlers

import (
	"context"
	"log"
	"test/todo/database"
	"test/todo/models"
	"test/todo/templates"

	"github.com/labstack/echo/v4"
)

func CreateTodo(ctx echo.Context) error {
    db, err := database.OpenDatabase()
    msg := "an error occured when trying to insert a todo: %s"

    if err != nil {
        log.Printf(msg, err)
        return renderServerError(ctx, err.Error())
    }
    
    store := database.NewTodoStore(db)

    var td models.Todo

    if err = ctx.Bind(&td); err != nil {
        return renderValidationError(ctx, err.Error())  
    }

    if err = store.CreateTodo(&td); err != nil {
        log.Printf(msg, err)
        return renderServerError(ctx, err.Error())
    }

    component := templates.TodoCreate("Create a todo" , td)

    return component.Render(context.Background(), ctx.Response().Writer)
} 
