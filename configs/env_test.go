package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEnv(t *testing.T) {
	t.Run("should return a new Env struct", func(t *testing.T) {
		env := newEnv()

		assert.NotNil(t, env)
		assert.IsType(t, &Env{}, env)
	})

	t.Run("should return a new Env struct with default values", func(t *testing.T) {
		env := newEnv()

		env.Port = DefaultPort

		assert.Equal(t, "3333", env.Port)
	})
}

func TestLoadEnv(t *testing.T) {
	t.Run("should load the .env file", func(t *testing.T) {
		LoadEnv("../.env")

		assert.NotNil(t, ENV)
		assert.IsType(t, &Env{}, ENV)
	})
}
