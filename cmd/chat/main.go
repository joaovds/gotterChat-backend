package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joaovds/chat/configs"
)

func main() {
  configs.LoadEnv()

  app := fiber.New()
  fmt.Println("Server running on port", configs.ENV.Port)
  log.Fatal(app.Listen(":" + configs.ENV.Port))
}
