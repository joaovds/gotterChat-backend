package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joaovds/chat/internal/infra/webserver/handlers"
)

func handleWebsocketRoutes(mux *chi.Mux) {
	websocketHandlers := handlers.NewWebsocketHandler()

	mux.Route("/ws", func(wsRouter chi.Router) {
		wsRouter.Get("/{roomId}", func(w http.ResponseWriter, r *http.Request) {
			roomId := chi.URLParam(r, "roomId")

			websocketHandlers.ServeWs(roomId, w, r)
		})
	})
}
