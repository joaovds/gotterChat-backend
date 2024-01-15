package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	apiV1 := app.Group("/api/v1")

  handleWebsocketRoutes(apiV1)
}
