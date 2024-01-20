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

const (
  DefaultPort = "3333"
)

func newEnv() *Env {
	port := "3333"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	return &Env{
		Port: port,
	}
}

func LoadEnv(filepath ...string) {
  if len(filepath) > 1 {
    log.Panic("You can only pass one filepath to LoadEnv")
  }

  if len(filepath) == 0 {
    filepath = append(filepath, ".env")
  }

	err := godotenv.Load(filepath[0])
	if err != nil {
		log.Panic("Error loading .env file:", err)
	}

	ENV = newEnv()
}
