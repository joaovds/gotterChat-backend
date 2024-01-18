package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Port       string
	MongoDBUri string
}

var ENV *Env

func newEnv() *Env {
	port := "3333"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	mongoDBURI := "mongodb://localhost:27017"
	if os.Getenv("MONGODB_URI") != "" {
		mongoDBURI = os.Getenv("MONGODB_URI")
	}

	return &Env{
		Port:       port,
		MongoDBUri: mongoDBURI,
	}
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file:", err)
	}

	ENV = newEnv()
}
