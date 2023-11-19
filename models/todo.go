package models

import (
    "fmt"
    "database/sql"
)

type Todo struct {
    Id           int       `db:"id"`
    Title        string    `db:"title" form:"title"`
	Content      sql.NullString `db:"content" form:"content"`
    Status       string    `db:"status" form:"status"`
	Priority     string    `db:"priority" form:"priority"`
}

func (td *Todo) String() string{
    return fmt.Sprintf("Todo {Title: %s, Status: %s, Priority: %s}",  td.Title, td.Status, td.Priority) 
}

