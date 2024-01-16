package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joaovds/chat/infra/webserver/handlers"
)

func handleWebsocketRoutes(mux *chi.Mux) {
	websocketHandlers := handlers.NewWebsocketHandler()

	mux.Route("/ws", func(wsRouter chi.Router) {
		wsRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
			websocketHandlers.ServeWs(w, r)
		})
	})
}
