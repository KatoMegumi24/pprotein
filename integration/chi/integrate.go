package chi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kaz/pprotein/integration"
)

func Integrate(r chi.Router) {
	EnableDebugHandler(r)
	EnableDebugMode(r)
}

func EnableDebugHandler(r chi.Router) {
	r.Route("/debug", func(r chi.Router) {
		r.Use(func(h http.Handler) http.Handler { return integration.NewDebugHandler() })
	})
}

func EnableDebugMode(r chi.Router) {
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
}
