package session

import (
	"fmt"
	"net/http"
	"time"
)

func SetAuthCookie(w http.ResponseWriter, login string, role string) {
	value := fmt.Sprintf("%s|%s", login, role)

	http.SetCookie(w, &http.Cookie{
		Name:     "auth",
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   28800,
		Expires:  time.Now().Add(8 * time.Hour),
		SameSite: http.SameSiteLaxMode,
		// Secure: true
	})
}

func ClearAuthCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "auth",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
		SameSite: http.SameSiteLaxMode,
	})
}

func IsAuthenticated(r *http.Request) bool {
	cookie, err := r.Cookie("auth")
	if err != nil {
		return false
	}
	return cookie.Value == "true"
}
