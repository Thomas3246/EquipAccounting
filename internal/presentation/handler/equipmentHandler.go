package handler

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/Thomas3246/EquipAccounting/internal/application/service"
	"github.com/Thomas3246/EquipAccounting/internal/domain"
	"github.com/Thomas3246/EquipAccounting/pkg/templateloader"
	"github.com/go-chi/chi/v5"
)

type EquipmentHandler struct {
	EquipService              service.EquipmentService
	EquipmentDirectoryService service.EquipmentDirectoryService
	DepartmentService         service.DepartmentService
}

func NewEquipmentHandler(
	equipService *service.EquipmentService,
	equipmentDirectoryService *service.EquipmentDirectoryService,
	depService *service.DepartmentService,
) *EquipmentHandler {
	return &EquipmentHandler{
		EquipService:              *equipService,
		EquipmentDirectoryService: *equipmentDirectoryService,
		DepartmentService:         *depService,
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

func (h *EquipmentHandler) EquipmentGet(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templateloader.GetTemplatePath("base.html"), templateloader.GetTemplatePath("equipment.html"))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Error by parse", err)
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	equipment, err := h.EquipService.GetEquipmentById(id)
	if err != nil {
		if err == service.ErrNotFound {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении оборудования: ", err)
		return
	}

	directories, err := h.EquipmentDirectoryService.GetEquipmentDirectoriesViewByFilter(0)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении оборудования: ", err)
		return
	}

	departments, err := h.DepartmentService.GetDepartmentsView()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении оборудования: ", err)
		return
	}

	states, err := h.EquipService.GetEquipmentStates()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении оборудования: ", err)
		return
	}

	templData := struct {
		IsAdmin     int
		Equipment   domain.Equipment
		Directories []domain.EquipmentDirectoryView
		Departments []domain.DepartmentView
		States      []domain.EquipmentState
		InvNum      string
		InvNumError string
	}{
		IsAdmin:     1,
		Equipment:   equipment,
		Directories: directories,
		Departments: departments,
		States:      states,
		InvNum:      "",
		InvNumError: "",
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, templData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error by execute ", err)
	}
}

func (h *EquipmentHandler) EquipmentPost(w http.ResponseWriter, r *http.Request) {
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

	invNum := r.Form.Get("invNum")
	directoryStr := r.Form.Get("directoryId")
	directory, _ := strconv.Atoi(directoryStr)
	departmentStr := r.Form.Get("departmentId")
	department, _ := strconv.Atoi(departmentStr)

	invNumIsFree, err := h.EquipService.CheckInvNumForFree(id, invNum)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при проверке инв. номера: ", err)
		return
	}

	equipment := domain.Equipment{
		Id:           id,
		InvNum:       invNum,
		DirectoryId:  directory,
		DepartmentId: department,
	}

	if invNumIsFree {
		err = h.EquipService.UpdateEquipment(equipment)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("Ошибка при изменении админом оборудования: ", err)
			return
		}
		http.Redirect(w, r, "/equipment", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles(templateloader.GetTemplatePath("base.html"), templateloader.GetTemplatePath("equipment.html"))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Error by parse", err)
		return
	}

	directories, err := h.EquipmentDirectoryService.GetEquipmentDirectoriesViewByFilter(0)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении оборудования: ", err)
		return
	}

	departments, err := h.DepartmentService.GetDepartmentsView()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении оборудования: ", err)
		return
	}

	states, err := h.EquipService.GetEquipmentStates()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении оборудования: ", err)
		return
	}

	oldEquipment, err := h.EquipService.GetEquipmentById(id)
	if err != nil {
		if err == service.ErrNotFound {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении оборудования: ", err)
		return
	}

	templData := struct {
		IsAdmin     int
		Equipment   domain.Equipment
		Directories []domain.EquipmentDirectoryView
		Departments []domain.DepartmentView
		States      []domain.EquipmentState
		InvNum      string
		InvNumError string
	}{
		IsAdmin:     1,
		Equipment:   oldEquipment,
		Directories: directories,
		Departments: departments,
		States:      states,
		InvNum:      equipment.InvNum,
		InvNumError: "Инвентарный номер занят",
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, templData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error by execute ", err)
	}

}

func (h *EquipmentHandler) DeleteEquipment(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = h.EquipService.DeleteEquipment(id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при удалении оборудования: ", err)
		return
	}

	http.Redirect(w, r, "/equipment", http.StatusSeeOther)
}
