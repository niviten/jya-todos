package todoAPI

import (
	"errors"
	"jya-todos/src/components/todo"
	"jya-todos/src/utils"

	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	type RequestData struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		DueBy       string `json:"dueBy"`
		Priority    int    `json:"priority"`
	}

	var data RequestData

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	t := todo.NewTodo()

	if len(data.Title) == 0 {
		return errors.New("Title is required")
	}
	t.Title = data.Title

	if len(data.DueBy) == 0 {
		t.DueBy = 0
	} else {
		dueBy, err := utils.ParseDateToEpoch(data.DueBy)
		if err != nil {
			return err
		}
		t.DueBy = dueBy
	}

	t.Description = data.Description

	if data.Priority <= 0 {
		t.Priority = 5
	} else {
		t.Priority = int8(data.Priority)
	}

	t.CreatedAt = utils.GetCurrentTimeInEpoch()

	lastInsertedId, err := todo.Create(t)
	if err != nil {
		return err
	}

	err = c.JSON(struct {
		LastInsertedId int    `json:"lastInsertedId"`
		Message        string `json:"message"`
	}{
		LastInsertedId: int(lastInsertedId),
		Message:        "Todo created successfully",
	})

	return err
}
