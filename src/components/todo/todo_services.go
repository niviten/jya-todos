package todo

import (
	"context"
	"jya-todos/src/connection"
)

func Create(todo *Todo) (int64, error) {
	db, err := connection.GetDBConnection()
	if err != nil {
		return -1, err
	}
	query := `
    INSERT INTO Todo (
        todo_title, todo_description, created_at, due_by, priority
    ) VALUES (?, ?, ?, ?, ?)
    `
	ctx := context.TODO()
	result, err := db.ExecContext(
		ctx,
		query,
		todo.Title,
		todo.Description,
		todo.CreatedAt,
		todo.DueBy,
		todo.Priority,
	)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}
