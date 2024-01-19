package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/joaovds/chat/configs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func handleDocsRoutes(mux *chi.Mux) {
	mux.Route("/docs", func(wsRouter chi.Router) {
		wsRouter.Get("/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:"+configs.ENV.Port+"/api/v1/docs/doc.json")))
	})
}
