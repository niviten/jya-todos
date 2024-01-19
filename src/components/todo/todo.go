package todo

type Todo struct {
    Id int64
    Title string
    Description string
    IsCompleted bool
    CreatedAt int64
    DueBy int64
    Priority int8
}

func NewTodo() *Todo {
    return &Todo{}
}
