package models

import "fmt"

type Todo struct {
    Id           int       `db:"id"`
    Title        string    `db:"title" form:"title"`
	Content      string    `db:"content" form:"content"`
    Status       string    `db:"status" form:"status"`
	Priority     string    `db:"priority" form:"priority"`
}

func (td *Todo) String() string{
    return fmt.Sprintf("Todo {Id: %d, Title: %s, Content: %s, Status: %s, Priority: %s}", td.Id, td.Title, td.Content, td.Status, td.Priority) 
}

