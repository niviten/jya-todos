package main

import (
	"fmt"
	"jya-todos/src/components/todo"
	"jya-todos/src/utils"
)

func main() {
    t := todo.NewTodo()
    t.Title = "Write code"
    t.Description = "Write go code to work on that project"
    t.CreatedAt = utils.GetCurrentTimeInEpoch()
    t.DueBy = utils.GetCurrentTimeInEpoch() + (2 * 24 * 60 * 60 * 1000)
    t.Priority = 1

    todoId, err := todo.Create(t)
    if err != nil {
        fmt.Println("Error occurred while creating todo", err)
        return
    }

    fmt.Println("Successfully todo created:", todoId)

    fmt.Println("___done")
}
