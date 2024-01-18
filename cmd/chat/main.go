package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joaovds/chat/application/repository"
	"github.com/joaovds/chat/configs"
	"github.com/joaovds/chat/infra/webserver/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	configs.LoadEnv()

	//db connection
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(configs.ENV.MongoDBUri))
	if err != nil {
		fmt.Println("Error connecting to MongoDB: ", err)
	}
	defer client.Disconnect(context.TODO())

	fmt.Println("MongoDB succesfully connected.")
	db := client.Database("gotter-chat")

	//repositories
	userRepository := repository.NewUserRepository(db)

	router := chi.NewRouter()

	routes.SetupRoutes(router, userRepository)

	fmt.Println("Server running on port", configs.ENV.Port)
	log.Fatal(http.ListenAndServe(":"+configs.ENV.Port, router))
}
