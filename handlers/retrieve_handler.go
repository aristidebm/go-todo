package handlers

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"

	"test/todo/models"
	"test/todo/templates"
)

func RetrieveTodo(c echo.Context) error {
   
    db, err := GetBD()

    if err != nil {
       return renderServerError(c, err.Error()) 
    }
    var id int 

    // creates query params binder that stops binding at first error
    // https://echo.labstack.com/docs/binding#fluent-binding
    err = echo.QueryParamsBinder(c). Int("id", &id).BindError()

    if err != nil {
        return renderValidationError(c, err.Error())  
    }

    store := NewTodoStore(db)

    var td models.Todo

    if td, err = store.RetrieveTodo(id); err != nil {
        return renderServerError(c, err.Error())  
    }

    component := templates.TodoRetrieve(fmt.Sprintf("Todo %d", td.Id), td)
    return component.Render(context.Background(), c.Response().Writer)
}
