package handlers

import ( 
    "context"
    "github.com/labstack/echo/v4"

    "test/todo/models"
    "test/todo/templates"
)

func ListTodo(c echo.Context) error {
    db, err := GetBD()

    if err != nil {
        return renderServerError(c, err.Error()) 
    }
    
    store := NewTodoStore(db)

    var todos []models.Todo 

    if todos, err = store.ListTodos(); err != nil {
        return renderServerError(c, err.Error()) 
    }
    component := templates.TodoList("List of todos" , todos)
    return component.Render(context.Background(), c.Response().Writer)
}
