package presentation

import (
	"net/http"

	"github.com/Thomas3246/EquipAccounting/internal/presentation/handler"
	"github.com/Thomas3246/EquipAccounting/internal/presentation/middleware"
	"github.com/go-chi/chi/v5"
	basicMW "github.com/go-chi/chi/v5/middleware"
)

func NewRouter(*handler.AppHandler) *chi.Mux {
	r := chi.NewMux()

	// global mw
	r.Use(middleware.LoggingMiddleware)
	r.Use(basicMW.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return r
}
