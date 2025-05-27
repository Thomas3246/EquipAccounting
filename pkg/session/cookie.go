package session

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func SetAuthCookie(w http.ResponseWriter, login string, isAdmin int) {
	value := fmt.Sprintf("%s|%d", login, isAdmin)
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

func GetIsAdminFromCookie(cookieValue string) (bool, bool) {

	parts := strings.Split(cookieValue, "|")
	if len(parts) != 2 {
		return false, false
	}

	isAdmin, err := strconv.Atoi(parts[1])
	if err != nil {
		return false, false
	}

	return isAdmin != 0, true
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
