package todoAPI

import (
	"github.com/gofiber/fiber/v2"
)

func AddRoutes(router fiber.Router) {
    r := router.Group("/todo")
    r.Post("/create", create)
}
