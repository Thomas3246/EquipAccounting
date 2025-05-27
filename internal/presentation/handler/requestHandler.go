package handler

import (
	"log"
	"net/http"
	"text/template"

	"github.com/Thomas3246/EquipAccounting/internal/application/service"
	"github.com/Thomas3246/EquipAccounting/pkg/templateloader"
)

type RequestHandler struct {
	service service.RequestService
}

func NewRequestHandler(service *service.RequestService) *RequestHandler {
	return &RequestHandler{service: *service}
}

func (h *RequestHandler) AllActiveRequests(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templateloader.GetTemplatePath("base.html"), templateloader.GetTemplatePath("allactive.html"))
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

	requests, err := h.service.GetAllActive(cookie.Value)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при извлечении заявок: ", err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, requests)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Println("Template Execute Error: ", err)
	}
}
