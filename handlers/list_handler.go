package handlers

import ( 
    "context"
    "github.com/labstack/echo/v4"
	// "github.com/a-h/templ"

    "test/todo/models"
    "test/todo/templates"
)

func ListTodo(c echo.Context) error {
    title := "Todo"
	todos := []models.Todo{
		{Title: "Touch Typing", Content: "My Learning touch typing with calmack-dh keybord", Priority: "High"},
		{Title: "English", Content: "I need to learn english", Priority: "Medium"},
		{Title: "French", Content: "I fluently speak french", Priority: "Low"},
	}
    component := templates.TodoList(title, todos)
    return component.Render(context.Background(), c.Response().Writer)
}
