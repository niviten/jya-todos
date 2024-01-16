package todo

import (
	"context"
	"fmt"
	"jya-todos/src/connection"
)

func Create(conn *connection.RedisConnection, todo *Todo) error {
    ctx := context.Background()
    var todoId int64 
    var err error
    todoId, err = conn.Client.Incr(ctx, "todo_id").Result()
    if err != nil {
        return err
    }
    todo.Id = todoId
    fmt.Println("todo:", todo)
    key := fmt.Sprintf("%s_%d", "todo", todo.Id)
    value := todo.String()
    err = conn.Client.Set(ctx, key, value, 0).Err()
    return err
}
