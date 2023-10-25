package main

import (
	"html/template"
	"net/http"
    "io"

	"github.com/labstack/echo/v4"
)

type Template struct {
   templates *template.Template 
}

type Todo struct {
    Title string
    Content string
    Priority string
}

// implement echo.Renderer interface 
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
   return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
    router := echo.New()
    renderer := Template{
        templates: template.Must(template.ParseFiles("todo.html")), 
    }
    router.Renderer = &renderer
    router.GET("/todos", func(c echo.Context) error {
        todos := []Todo {
            {Title: "Touch Typing", Content: "My Learning touch typing with calmack-dh keybord", Priority: "High"},
            {Title: "English", Content: "I need to learn english", Priority: "Medium"},
            {Title: "French", Content: "I fluently speak french", Priority: "Low"},
        }
        return c.Render(http.StatusOK, "todos", todos)
    })
    router.Logger.Fatal(router.Start(":8000"))
}
