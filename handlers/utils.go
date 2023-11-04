package handlers

import (
    "fmt"
    "context"
    "github.com/labstack/echo/v4"
	"github.com/a-h/templ"
	"github.com/jmoiron/sqlx"

	"test/todo/models"
   "test/todo/templates"
)

func Render(c echo.Context, t templ.Component) error{
    c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
    return t.Render(c.Request().Context(), c.Response().Writer)
}

func GetBD() (*sqlx.DB, error) {
    // exactly the same as the built-in
    db, err := sqlx.Open("sqlite3", ":memory:")
    // force a connection and test that it worked
    if err != nil {
        return nil, err
    }
    // check that the database if reachable.
    if err = db.Ping(); err != nil {
        return nil, err 
    }
    return db, nil
}

type TodoStore struct {
    *sqlx.DB    
}

func (s *TodoStore) CreateTodo(td *models.Todo) error {
    tx, err := s.Beginx()

    if err != nil {
        return fmt.Errorf("Cannot insert this row %s", td) 
    }
    
    _, err = tx.Exec(`INSERT INTO todos(title, content, status, priority) VALUES ($1, $2, $3, $4);`, 
        td.Title,
        td.Content,
        td.Status,
        td.Priority)

    if err != nil {
        //  FIXME: Add proper error handling
        return fmt.Errorf("Cannot insert this row %s", td) 
    }

    var id int

    err = tx.Get(id, `SELECT LAST_INSERT_ROWID();`)

    if err != nil {
        //  FIXME: Add proper error handling
        return fmt.Errorf("Cannot insert this row %s", td) 
    }

    td.Id = id

    tx.Commit()

    return nil
}

func (s *TodoStore) ListTodos() ([]models.Todo, error) {
    tds := []models.Todo{}
    if err := s.Select(tds, `SELECT * FROM todos;`); err != nil {
        return tds, fmt.Errorf("Cannot list todos: %w ", err) 
    }
    return tds, nil
}

func (s *TodoStore) RetrieveTodo(id int) (models.Todo, error) {
    var td models.Todo 
    if err := s.Get(&td, `SELECT * FROM todos WHERE id = $1;`, id); err != nil {
        return td, fmt.Errorf("Cannot find a todo with id = %d: %w", id, err) 
    }
    return td, nil
}

func (s *TodoStore) UpdateTodo(td *models.Todo) error { 
    _, err := s.Exec(`UPDATE todos SET title = $1, content = $2, status = $3 priority = $4 WHERE id = $5;`,
    td.Title,
    td.Content,
    td.Status,
    td.Priority,
    td.Id)
    
    if err != nil {
        return fmt.Errorf("Cannot find a todo with id = %d: %w", td.Id, err) 
    }
    return nil 
}

func (s *TodoStore) DeleteTodo(id int) error {
    if _, err := s.Exec(`DELETE FROM todos WHERE id = $1;`, id); err != nil {
        return fmt.Errorf("Cannot find a todo with id = %d: %w", id, err) 
    }
    return nil 
}

func NewTodoStore(db *sqlx.DB) (s *TodoStore) {
    return &TodoStore{ db }
}

func renderServerError(c echo.Context, stackTrace string) error {
    component := templates.ServerError("500 Internal Error" , stackTrace)
    return component.Render(context.Background(), c.Response().Writer)
}

func renderValidationError(c echo.Context, message string) error {
    component := templates.ValidationError("400 Bad Request" , message)
    return component.Render(context.Background(), c.Response().Writer)
}

func renderNotFoundError(c echo.Context, message string) error {
    component := templates.ValidationError("404 Not found" , message)
    return component.Render(context.Background(), c.Response().Writer)
}

