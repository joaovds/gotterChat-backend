package routes

import (
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(mainMux *chi.Mux) {
	apiV1 := chi.NewRouter()

	handleWebsocketRoutes(apiV1)
	handleRoomRoutes(apiV1)
	handleDocsRoutes(apiV1)
	handleUserRoutes(apiV1)

	mainMux.Mount("/api/v1", apiV1)
}
