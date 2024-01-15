package routes

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func handleWebsocketRoutes(router fiber.Router) {
  wsRouter := router.Group("/ws") 

  wsRouter.Use("/", func (c *fiber.Ctx) error {
    if websocket.IsWebSocketUpgrade(c) {
      c.Locals("allowed", true)
      return c.Next()
    }
    return fiber.ErrUpgradeRequired
  })

  wsRouter.Get("/", websocket.New(serveHs))
}

func serveHs(c *websocket.Conn) {
  for {
    mt, msg, err := c.ReadMessage()
    if err != nil {
      log.Println("read:", err)
      break
    }

    log.Printf("recv: %s", msg)

    err = c.WriteMessage(mt, msg)
    if err != nil {
      log.Println("write:", err)
      break
    }
  }
}
