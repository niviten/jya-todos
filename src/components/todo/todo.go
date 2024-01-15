package todo

import (
	"strconv"
	"strings"
)

type Todo struct {
    Id int64
    Name string
    Description string
    IsCompleted bool
    DueBy int64
    Priority int8
}

func NewTodo() *Todo {
    return &Todo{-1, "", "", false, -1, 5}
}

func (t *Todo) String() string {
    fields := []string{
        strconv.FormatInt(t.Id, 10),
        t.Name,
        t.Description,
        strconv.FormatBool(t.IsCompleted),
        strconv.FormatInt(t.DueBy, 10),
        strconv.FormatInt(int64(t.Priority), 10),
    }
    return strings.Join(fields, "::")
}
