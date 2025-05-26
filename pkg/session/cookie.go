package session

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func SetAuthCookie(w http.ResponseWriter, login string, departmentId int, isAdmin int) {
	value := fmt.Sprintf("%s|%d|%d", login, departmentId, isAdmin)
	fmt.Print(value)

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

func GetUserRoleFromCookie(cookieValue string) (string, bool) {

	parts := strings.Split(cookieValue, "|")
	if len(parts) != 2 {
		return "", false
	}

	return parts[1], true
}

func IsAuthenticated(r *http.Request) bool {
	cookie, err := r.Cookie("auth")
	if err != nil {
		return false
	}

	parts := strings.Split(cookie.Value, "|")
	if len(parts) != 2 {
		return false
	}

	return parts[0] != "" && parts[1] != ""
}
