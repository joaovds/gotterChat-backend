package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joaovds/chat/application/repository"
	"github.com/joaovds/chat/configs"
	"github.com/joaovds/chat/internal/infra/database"
	"github.com/joaovds/chat/internal/infra/webserver/routes"
)

func main() {
	configs.LoadEnv()

	mongoConfig := database.MongoDBConfig{
		URI:      configs.ENV.MongoDbUri,
		Database: configs.ENV.MongoDbName,
	}

	mongoInstance := database.SetupMongoDB(mongoConfig)
	defer mongoInstance.Client.Disconnect(context.TODO())

	userRepository := repository.NewUserRepository(mongoInstance.Db)

	router := chi.NewRouter()

	routes.SetupRoutes(router, userRepository)

	fmt.Println("Server running on port", configs.ENV.Port)
	log.Fatal(http.ListenAndServe(":"+configs.ENV.Port, router))
}
