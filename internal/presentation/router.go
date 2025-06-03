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

		r.Get("/allactive", h.RequestHandler.AllActiveRequests)
		r.Get("/allactive/{login}", h.RequestHandler.AllActiveUserRequests)

		r.Get("/allclosed", h.RequestHandler.AllClosedRequests)
		r.Get("/allclosed/{login}", h.RequestHandler.AllClosedUserRequests)

		r.Get("/newRequest", h.RequestHandler.NewRequestGet)
		r.Post("/newRequest", h.RequestHandler.NewRequestPost)

		r.Get("/request/{id}", h.RequestHandler.RequestEditGet)
		r.Post("/request/{id}", h.RequestHandler.RequestEditPost)

		// admin only
		r.Group(func(r chi.Router) {
			r.Use(middleware.AdminMiddleWare)

			r.Get("/users", h.UserHandler.Users)

			r.Get("/users/new", h.UserHandler.AddUserGet)
			r.Post("/users/new", h.UserHandler.AddUserPost)

			r.Get("/users/{id}", h.UserHandler.UserGet)
			r.Post("/users/{id}", h.UserHandler.UserPost)

			r.Post("/users/{id}/delete", h.UserHandler.DeleteUser)

		})

	})

	return r
}
