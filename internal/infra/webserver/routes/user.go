package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/joaovds/chat/internal/infra/webserver/handlers"
)

func handleUserRoutes(mux *chi.Mux) {
	userHandlers := handlers.NewUserHandler()

	mux.Route("/users", func(wsRouter chi.Router) {
		wsRouter.Post("/", userHandlers.CreateUser)
	})
}
