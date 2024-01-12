package todo

import (
	"fmt"
)

type Todo struct {
    Id int64
    Name string
    Description string
    IsCompleted bool
    DueBy int64
}

func (todo *Todo) ToString() string {
    var isCompleted int
    if todo.IsCompleted {
        isCompleted = 1
    } else {
        isCompleted = 0
    }
    return fmt.Sprintf("%ld::%s::%s::%d::%ld", todo.Id, todo.Name,
        todo.Description, isCompleted, todo.DueBy)
}
