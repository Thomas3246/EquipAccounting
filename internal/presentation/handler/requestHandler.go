package handler

import (
	"net/http"

	"github.com/Thomas3246/EquipAccounting/internal/application/service"
)

type RequestHandler struct {
	service service.RequestService
}

func NewRequestHandler(service *service.RequestService) *RequestHandler {
	return &RequestHandler{service: service}
}

func (h *RequestHandler) AllActiveRequests(w http.ResponseWriter, r *http.Request) {

	добавить функционал

}
