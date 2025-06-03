package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/Thomas3246/EquipAccounting/internal/application/service"
	"github.com/Thomas3246/EquipAccounting/internal/domain"
	"github.com/Thomas3246/EquipAccounting/pkg/session"
	"github.com/Thomas3246/EquipAccounting/pkg/templateloader"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	userService       service.UserService
	departmentService service.DepartmentService
}

func NewUserHandler(
	userService *service.UserService,
	departmentService *service.DepartmentService,
) *UserHandler {
	return &UserHandler{
		userService:       *userService,
		departmentService: *departmentService,
	}
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

	user, err := h.userService.Authenticate(login, password)
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

	session.SetAuthCookie(w, user.Login, user.IsAdmin)

	http.Redirect(w, r, "/allactive", http.StatusSeeOther)
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	session.ClearAuthCookie(w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (h *UserHandler) Dashboard(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templateloader.GetTemplatePath("base.html"), templateloader.GetTemplatePath("dashboard.html"))
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
	name := query.Get("name")
	isAdmin := query.Get("isAdmin")
	department := query.Get("department")

	err := h.userService.Register(login, password, name, isAdmin, department)
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

func (h *UserHandler) Users(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templateloader.GetTemplatePath("base.html"), templateloader.GetTemplatePath("users.html"))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Error by parse", err)
		return
	}

	users, err := h.userService.GetUsers()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении пользователей: ", err)
		return
	}

	cookie, err := r.Cookie("auth")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при извлечении cookie: ", err)
		return
	}
	parts := strings.Split(cookie.Value, "|")
	if len(parts) != 2 {
		http.Error(w, "Invalid Cookie Value", http.StatusInternalServerError)
		return
	}

	isAdmin, err := strconv.Atoi(parts[1])
	if err != nil {
		http.Error(w, "Invalid Cookie Value", http.StatusInternalServerError)
		log.Println("Ошибка чтения значения isAdmin в cookie: ", err)
		return
	}

	templData := struct {
		Users   []domain.ViewUser
		IsAdmin int
	}{
		Users:   users,
		IsAdmin: isAdmin,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, templData)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Println("Template Execute Error: ", err)
	}
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = h.userService.DeleteUser(userId)
	if err != nil {

		if err == service.ErrNoAccess {
			http.Error(w, "У вас нет доступа для этого действия", http.StatusForbidden)
			return
		}

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка удаления пользователя: ", err)
		return
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func (h *UserHandler) AddUserGet(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templateloader.GetTemplatePath("base.html"), templateloader.GetTemplatePath("newUser.html"))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Error by parse", err)
		return
	}

	cookie, err := r.Cookie("auth")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при извлечении cookie: ", err)
		return
	}
	parts := strings.Split(cookie.Value, "|")
	if len(parts) != 2 {
		http.Error(w, "Invalid Cookie Value", http.StatusInternalServerError)
		return
	}
	isAdmin, err := strconv.Atoi(parts[1])
	if err != nil {
		http.Error(w, "Invalid Cookie Value", http.StatusInternalServerError)
		log.Println("Ошибка чтения значения isAdmin в cookie: ", err)
		return
	}

	departments, err := h.departmentService.GetDepartmentsView()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка получения отделов: ", err)
		return
	}

	templData := struct {
		Departments []domain.DepartmentView
		IsAdmin     int
	}{
		Departments: departments,
		IsAdmin:     isAdmin,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, templData)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Println("Template Execute Error: ", err)
	}
}

func (h *UserHandler) AddUserPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	login := r.Form.Get("login")
	name := r.Form.Get("name")
	password := r.Form.Get("password")
	department := r.Form.Get("department_id")
	departmentId, _ := strconv.Atoi(department)

	user := domain.User{
		Login:        login,
		Name:         name,
		Password:     password,
		DepartmentId: departmentId,
	}

	err = h.userService.AddUser(user)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при добавлении пользователя: ", err)
		return
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func (h *UserHandler) UserGet(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templateloader.GetTemplatePath("base.html"), templateloader.GetTemplatePath("user.html"))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Error by parse", err)
		return
	}

	cookie, err := r.Cookie("auth")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при извлечении cookie: ", err)
		return
	}
	parts := strings.Split(cookie.Value, "|")
	if len(parts) != 2 {
		http.Error(w, "Invalid Cookie Value", http.StatusInternalServerError)
		return
	}
	isAdmin, err := strconv.Atoi(parts[1])
	if err != nil {
		http.Error(w, "Invalid Cookie Value", http.StatusInternalServerError)
		log.Println("Ошибка чтения значения isAdmin в cookie: ", err)
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	user, err := h.userService.GetUserById(id)
	if err != nil {
		if err == service.ErrUserNotFound {
			http.Error(w, "User Not Found", http.StatusNotFound)
			log.Println(err)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка получения пользователя: ", err)
		return
	}

	departments, err := h.departmentService.GetDepartmentsView()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка получения отделов: ", err)
		return
	}

	templData := struct {
		User        domain.User
		IsAdmin     int
		Departments []domain.DepartmentView
	}{
		User:        user,
		IsAdmin:     isAdmin,
		Departments: departments,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, templData)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Println("Template Execute Error: ", err)
	}
}

func (h *UserHandler) UserPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	login := r.Form.Get("login")
	name := r.Form.Get("name")
	password := r.Form.Get("password")
	departmentIdStr := r.Form.Get("department_id")
	departmentId, _ := strconv.Atoi(departmentIdStr)

	user := domain.User{
		Id:           id,
		Login:        login,
		Name:         name,
		Password:     password,
		DepartmentId: departmentId,
	}

	err = h.userService.EditUser(user)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при редактировании пользователя: ", err)
		return
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
