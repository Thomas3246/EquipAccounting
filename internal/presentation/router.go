package presentation

import (
	"github.com/Thomas3246/EquipAccounting/internal/presentation/handler"
	"github.com/Thomas3246/EquipAccounting/internal/presentation/middleware"
	"github.com/go-chi/chi/v5"
	basicMW "github.com/go-chi/chi/v5/middleware"
)

func NewRouter(h *handler.AppHandler) *chi.Mux {
	r := chi.NewMux()

	// global mw
	r.Use(middleware.LoggingMiddleware)
	r.Use(basicMW.Recoverer)

	// non auth
	r.Group(func(r chi.Router) {
		r.Get("/login", h.UserHandler.LoginGet)
		r.Post("/login", h.UserHandler.LoginPost)
		r.Get("/logout", h.UserHandler.Logout)

	})

	// auth
	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMidlleWare)
		r.Get("/dashboard", h.UserHandler.Dashboard)
		r.Get("/allactive", h.RequestHandler.AllActiveRequests)
		r.Get("/allactive/{login}", h.RequestHandler.AllActiveUserRequests)

		// admin only
		r.Group(func(r chi.Router) {
			r.Use(middleware.AdminMiddleWare)
			r.Get("/register", h.UserHandler.Register)
		})

	})

	return r
}
