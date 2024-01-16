package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joaovds/chat/configs"
	"github.com/joaovds/chat/infra/webserver/routes"
)

func main() {
	configs.LoadEnv()

	router := chi.NewRouter()

	routes.SetupRoutes(router)

	fmt.Println("Server running on port", configs.ENV.Port)
	log.Fatal(http.ListenAndServe(":"+configs.ENV.Port, router))
}
