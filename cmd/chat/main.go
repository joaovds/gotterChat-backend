package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joaovds/chat/configs"
	_ "github.com/joaovds/chat/docs"
	"github.com/joaovds/chat/infra/webserver/routes"
)

// @title Gotter Chat
// @version 1.0
// @description Chat Realtime with WebSocket, Golang and MongoDb (Front-end: Flutter)

// @license.name MIT
// @license.url https://github.com/joaovds/gotterChat-backend/blob/main/LICENSE

// @host localhost:3333
// @BasePath /api/v1
func main() {
	configs.LoadEnv()

	router := chi.NewRouter()

	routes.SetupRoutes(router)

	fmt.Println("Server running on port", configs.ENV.Port)
	log.Fatal(http.ListenAndServe(":"+configs.ENV.Port, router))
}
