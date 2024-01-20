package todo

import (
	"fmt"
	"time"
)

type Todo struct {
	Id          int64
	Title       string
	Description string
	IsCompleted bool
	CreatedAt   int64
	DueBy       int64
	Priority    int8
}

func NewTodo() *Todo {
	return &Todo{}
}

func (t *Todo) String() string {
	completedStatus := "Incomplete"
	if t.IsCompleted {
		completedStatus = "Completed"
	}

	createdAtStr := time.Unix(t.CreatedAt, 0).Format("2006-01-02 15:04:05")
	dueByStr := time.Unix(t.DueBy, 0).Format("2006-01-02 15:04:05")

	return fmt.Sprintf("Todo ID: %d\nTitle: %s\nDescription: %s\nStatus: %s\nCreated At: %s\nDue By: %s\nPriority: %d",
		t.Id, t.Title, t.Description, completedStatus, createdAtStr, dueByStr, t.Priority)
}
