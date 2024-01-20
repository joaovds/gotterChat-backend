package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Port        string
	MongoDbUri  string
	MongoDbName string
}

var ENV *Env

func newEnv() *Env {
	port := "3333"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	mongoDbURI := "mongodb://localhost:27017"
	if os.Getenv("MONGODB_URI") != "" {
		mongoDbURI = os.Getenv("MONGODB_URI")
	}

	mongoDbName := ""
	if os.Getenv("MONGO_DB_NAME") != "" {
		mongoDbName = os.Getenv("MONGO_DB_NAME")
	}

	return &Env{
		Port:        port,
		MongoDbUri:  mongoDbURI,
		MongoDbName: mongoDbName,
	}
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file:", err)
	}

	ENV = newEnv()
}
