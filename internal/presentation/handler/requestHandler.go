package handler

import (
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/Thomas3246/EquipAccounting/internal/application/service"
	"github.com/Thomas3246/EquipAccounting/internal/domain"
	"github.com/Thomas3246/EquipAccounting/pkg/templateloader"
	"github.com/go-chi/chi/v5"
)

type TemplateData struct {
	Requests  []domain.RequestView
	UserLogin string
	Flag      string
}

type RequestHandler struct {
	reqService  service.RequestService
	userService service.UserService
	eqService   service.EquipmentService
}

func NewRequestHandler(reqService *service.RequestService, userService *service.UserService, eqService *service.EquipmentService) *RequestHandler {
	return &RequestHandler{
		reqService:  *reqService,
		userService: *userService,
		eqService:   *eqService,
	}
}

func (h *RequestHandler) AllActiveRequests(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templateloader.GetTemplatePath("base.html"), templateloader.GetTemplatePath("requests.html"))
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Println("Template Parse Error: ", err)
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

	requests, err := h.reqService.GetAllActive(cookie.Value)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при извлечении заявок: ", err)
		return
	}

	templData := TemplateData{
		Requests:  requests,
		UserLogin: parts[0],
		Flag:      "allactive",
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, templData)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Println("Template Execute Error: ", err)
	}
}

func (h *RequestHandler) AllActiveUserRequests(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templateloader.GetTemplatePath("base.html"), templateloader.GetTemplatePath("requests.html"))
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Println("Template Parse Error: ", err)
		return
	}

	cookie, err := r.Cookie("auth")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при извлечении cookie: ", err)
		return
	}

	requestedLogin := chi.URLParam(r, "login")

	requests, err := h.reqService.GetAllUserActive(cookie.Value, requestedLogin)
	if err != nil {
		if err == service.ErrNoAccess {
			http.Error(w, "No Access To The Page", http.StatusForbidden)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при извлечении заявок: ", err)
		return
	}

	parts := strings.Split(cookie.Value, "|")
	if len(parts) != 2 {
		http.Error(w, "Invalid Cookie Value", http.StatusInternalServerError)
		return
	}
	templData := TemplateData{
		Requests:  requests,
		UserLogin: parts[0],
		Flag:      "alluseractive",
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, templData)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Println("Template Execute Error: ", err)
	}
}

func (h *RequestHandler) AllClosedRequests(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templateloader.GetTemplatePath("base.html"), templateloader.GetTemplatePath("requests.html"))
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Println("Template Parse Error: ", err)
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

	requests, err := h.reqService.GetAllClosed(cookie.Value)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при извлечении заявок: ", err)
		return
	}

	templData := TemplateData{
		Requests:  requests,
		UserLogin: parts[0],
		Flag:      "allclosed",
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, templData)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Println("Template Execute Error: ", err)
	}
}

func (h *RequestHandler) AllClosedUserRequests(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templateloader.GetTemplatePath("base.html"), templateloader.GetTemplatePath("requests.html"))
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Println("Template Parse Error: ", err)
		return
	}

	cookie, err := r.Cookie("auth")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при извлечении cookie: ", err)
		return
	}

	requestedLogin := chi.URLParam(r, "login")

	requests, err := h.reqService.GetAllUserClosed(cookie.Value, requestedLogin)
	if err != nil {
		if err == service.ErrNoAccess {
			http.Error(w, "No Access To The Page", http.StatusForbidden)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при извлечении заявок: ", err)
		return
	}

	parts := strings.Split(cookie.Value, "|")
	if len(parts) != 2 {
		http.Error(w, "Invalid Cookie Value", http.StatusInternalServerError)
		return
	}
	templData := TemplateData{
		Requests:  requests,
		UserLogin: parts[0],
		Flag:      "alluserclosed",
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, templData)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Println("Template Execute Error: ", err)
	}
}

func (h *RequestHandler) NewRequestGet(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templateloader.GetTemplatePath("base.html"), templateloader.GetTemplatePath("newRequest.html"))
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Println("Template Parse Error: ", err)
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

	requestTypes, err := h.reqService.GetRequestTypes()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении типов заявок: ", err)
		return
	}

	equipment, err := h.eqService.GetAvailableEquipment(parts[0])
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении доступного оборудования: ", err)
		return
	}

	templData := struct {
		UserLogin    string
		RequestTypes []domain.RequestType
		Equipment    []domain.EquipmentView
	}{
		UserLogin:    parts[0],
		RequestTypes: requestTypes,
		Equipment:    equipment,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, templData)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Println("Template Execute Error: ", err)
	}
}
