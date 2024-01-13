package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joaovds/chat/configs"
	"github.com/joaovds/chat/infra/fiber/routes"
)

func main() {
  configs.LoadEnv()

  app := fiber.New()

  routes.SetupRoutes(app)

  fmt.Println("Server running on port", configs.ENV.Port)
  log.Fatal(app.Listen(":" + configs.ENV.Port))
}
