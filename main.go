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
        name := c.QueryParam("name")
        if name == "" {
            name = "Mark"
        }
        return c.Render(http.StatusOK, "todos", name)
    })
    router.Logger.Fatal(router.Start(":8000"))
}
