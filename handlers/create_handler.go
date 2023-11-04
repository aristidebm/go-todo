package handlers

import ( 
    "context"
    "github.com/labstack/echo/v4"
   "test/todo/templates"
    "test/todo/models"
)

func CreateTodo(c echo.Context) error {
    db, err := GetBD()

    if err != nil {
        return renderServerError(c, err.Error())
    }
    
    store := NewTodoStore(db)

    var td models.Todo

    if err = c.Bind(&td); err != nil {
        return renderValidationError(c, err.Error())  
    }

    if err = store.CreateTodo(&td); err != nil {
        return renderServerError(c, err.Error())
    }

    component := templates.TodoCreate("Create a todo" , td)
    return component.Render(context.Background(), c.Response().Writer)
} 
