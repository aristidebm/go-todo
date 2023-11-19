package handlers

import (
	"context"
	"fmt"
    "log"

	"github.com/labstack/echo/v4"

	"test/todo/models"
	"test/todo/templates"
	"test/todo/database"
)

func RetrieveTodo(ctx echo.Context) error {
   
    db, err := database.OpenDatabase()

    if err != nil {
       return renderServerError(ctx, err.Error()) 
    }

    var id int 


    // creates query params binder that stops binding at first error
    // https://echo.labstack.com/docs/binding#fluent-binding
    err = echo.PathParamsBinder(ctx).Int("id", &id).BindError()

    if err != nil {
        return renderValidationError(ctx, err.Error())  
    }

    log.Printf("retrieving todo with id=%d", id) 

    store := database.NewTodoStore(db)

    var td models.Todo

    if td, err = store.RetrieveTodo(id); err != nil {
        return renderServerError(ctx, err.Error())  
    }

    component := templates.TodoRetrieve(fmt.Sprintf("Todo %d", td.Id), td)

    return component.Render(context.Background(), ctx.Response().Writer)
}
