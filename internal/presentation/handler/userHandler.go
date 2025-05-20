package handler

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/Thomas3246/EquipAccounting/internal/application/service"
	"github.com/Thomas3246/EquipAccounting/pkg/session"
	templateloader "github.com/Thomas3246/EquipAccounting/pkg/templateLoader"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: *service}
}

func (h *UserHandler) LoginGet(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(templateloader.GetTemplatePath("loginPage.html"))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Error by parse", err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error by execute")
	}
}

func (h *UserHandler) LoginPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	login := r.FormValue("login")
	password := r.FormValue("password")

	user, err := h.service.Authenticate(login, password)
	if err != nil {

		tmpl, err := template.ParseFiles(templateloader.GetTemplatePath("loginPage.html"))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("Template parse error: ", err)
			return
		}

		data := map[string]any{
			"Error": "Неверный логин или пароль",
			"Login": login,
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("Template Execute Error: ", err)
		}
		return
	}

	session.SetAuthCookie(w, user.Login, user.Role)

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	session.ClearAuthCookie(w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (h *UserHandler) Dashboard(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templateloader.GetTemplatePath("dashboard.html"))
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Println("Template Parse Error: ", err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Println("Template Execute Error: ", err)
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	login := query.Get("login")
	password := query.Get("password")
	role := query.Get("role")

	err := h.service.Register(login, password, role)
	if err != nil {
		if err == service.ErrNullParameter {
			http.Error(w, "Bad Request (missing parameter)", http.StatusBadRequest)
			return
		}

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при создании пользователя: ", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	resultStr := fmt.Sprintf("User created: %s %s", login, password)
	w.Write([]byte(resultStr))
}
