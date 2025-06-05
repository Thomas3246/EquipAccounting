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

type EquipmentDirectoryHandler struct {
	EquipmentDirectoryService service.EquipmentDirectoryService
	EquipmentService          service.EquipmentService
}

func NewEquipmentDirectoryHandler(eqDirService *service.EquipmentDirectoryService, eqService *service.EquipmentService) *EquipmentDirectoryHandler {
	return &EquipmentDirectoryHandler{EquipmentDirectoryService: *eqDirService, EquipmentService: *eqService}
}

func (h *EquipmentDirectoryHandler) DirectoryList(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templateloader.GetTemplatePath("base.html"), templateloader.GetTemplatePath("equipmentDirectoryList.html"))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Error by parse", err)
		return
	}

	eTypeStr := r.URL.Query().Get("type")
	var eType int
	if eTypeStr != "" {
		eType, _ = strconv.Atoi(eTypeStr)
	}

	eqTypes, err := h.EquipmentDirectoryService.GetEquipmentDirectoryTypes()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении типов оборудования: ", err)
		return
	}

	directories, err := h.EquipmentDirectoryService.GetEquipmentDirectoriesViewByFilter(eType)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении справочника оборудования: ", err)
		return
	}

	templData := struct {
		IsAdmin              int
		EquipmentTypes       []domain.EquipmentDirectoryType
		CurrentType          int
		EquipmentDirectories []domain.EquipmentDirectoryView
	}{
		IsAdmin:              1,
		EquipmentTypes:       eqTypes,
		CurrentType:          eType,
		EquipmentDirectories: directories,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, templData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error by execute ", err)
	}
}
