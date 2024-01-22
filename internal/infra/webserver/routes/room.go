package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joaovds/chat/internal/infra/webserver/handlers"
)

func handleRoomRoutes(mux *chi.Mux) {
	roomHandlers := handlers.NewRoomHandler()

  mux.Route("/rooms", func(wsRouter chi.Router) {
    wsRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
      roomHandlers.GetRooms(w, r)
    })
  })
}
