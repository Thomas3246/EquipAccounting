package handler

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/Thomas3246/EquipAccounting/internal/application/service"
	"github.com/Thomas3246/EquipAccounting/internal/domain"
	"github.com/Thomas3246/EquipAccounting/pkg/templateloader"
)

type EquipmentHandler struct {
	EquipService      service.EquipmentService
	DepartmentService service.DepartmentService
}

func NewEquipmentHandler(equipService *service.EquipmentService, depService *service.DepartmentService) *EquipmentHandler {
	return &EquipmentHandler{
		EquipService:      *equipService,
		DepartmentService: *depService,
	}
}

func (h *EquipmentHandler) EquipmentList(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templateloader.GetTemplatePath("base.html"), templateloader.GetTemplatePath("equipmentList.html"))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Error by parse", err)
		return
	}

	stateFilter := r.URL.Query().Get("state")
	departmentFilter := r.URL.Query().Get("department")
	var depID, stateID int
	if departmentFilter != "" {
		depID, _ = strconv.Atoi(departmentFilter)
	}
	if stateFilter != "" {
		stateID, _ = strconv.Atoi(stateFilter)
	}

	states, err := h.EquipService.GetEquipmentStates()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка получения статусов оборудования: ", err)
		return
	}

	departments, err := h.DepartmentService.GetDepartmentsView()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка получения отделов: ", err)
		return
	}

	equipment, err := h.EquipService.GetEquipmentViewByFilter(depID, stateID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка получения оборудования: ", err)
		return
	}

	templData := struct {
		IsAdmin           int
		Equipment         []domain.EquipmentView
		CurrentState      int
		CurrentDepartment int
		States            []domain.EquipmentState
		Departments       []domain.DepartmentView
	}{
		IsAdmin:           1,
		Equipment:         equipment,
		CurrentState:      stateID,
		CurrentDepartment: depID,
		States:            states,
		Departments:       departments,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, templData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error by execute ", err)
	}
}
