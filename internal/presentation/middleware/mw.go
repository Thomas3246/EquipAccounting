package middleware

import (
	"log"
	"net/http"

	"github.com/Thomas3246/EquipAccounting/pkg/session"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func AuthMidlleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !session.IsAuthenticated(r) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func AdminMiddleWare(nex http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		isAdmin, ok := session.GetIsAdminFromCookie(cookie.Value)
		if !ok {
			http.Error(w, "Invalid Session", http.StatusUnauthorized)
			return
		}

		if !isAdmin {
			http.Redirect(w, r, "/allactive", http.StatusSeeOther)
			return
		}

		nex.ServeHTTP(w, r)
	})
}
