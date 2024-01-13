package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	apiV1 := app.Group("/api/v1")

  apiV1.Get("/ping", func(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
      "message": "pong",
    })
  })
}
