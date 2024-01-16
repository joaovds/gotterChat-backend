package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Port string
}

var ENV *Env

func newEnv() *Env {
	port := "3333"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	return &Env{
		Port: port,
	}
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file:", err)
	}

	ENV = newEnv()
}
