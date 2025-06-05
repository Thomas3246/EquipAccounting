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

func (h *EquipmentDirectoryHandler) DirectoryGet(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	tmpl, err := template.ParseFiles(templateloader.GetTemplatePath("base.html"), templateloader.GetTemplatePath("directory.html"))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Error by parse", err)
		return
	}

	directory, err := h.EquipmentDirectoryService.GetEquipmentDirectory(id)
	if err != nil {
		if err == service.ErrRequestNotFound {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка получения справочника: ", err)
		return
	}

	types, err := h.EquipmentDirectoryService.GetEquipmentDirectoryTypes()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении типов оборудования: ", err)
		return
	}

	templData := struct {
		IsAdmin   int
		Directory domain.EquipmentDirectory
		Types     []domain.EquipmentDirectoryType
	}{
		IsAdmin:   1,
		Directory: directory,
		Types:     types,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, templData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error by execute ", err)
	}

}

func (h *EquipmentDirectoryHandler) DirectoryPost(w http.ResponseWriter, r *http.Request) {
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

	name := r.Form.Get("name")
	releaseYear := r.Form.Get("releaseYear")
	eType := r.Form.Get("typeId")
	typeId, _ := strconv.Atoi(eType)

	directory := domain.EquipmentDirectory{
		Id:          id,
		Name:        name,
		ReleaseYear: releaseYear,
		TypeId:      typeId,
	}

	err = h.EquipmentDirectoryService.UpdateDirectory(directory)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при изменении директории: ", err)
		return
	}

	http.Redirect(w, r, "/equipmentDirectory", http.StatusSeeOther)
}

func (h *EquipmentDirectoryHandler) NewDirectoryGet(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templateloader.GetTemplatePath("base.html"), templateloader.GetTemplatePath("newDirectory.html"))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Error by parse", err)
		return
	}

	types, err := h.EquipmentDirectoryService.GetEquipmentDirectoryTypes()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении типов оборудования: ", err)
		return
	}

	templData := struct {
		IsAdmin int
		Types   []domain.EquipmentDirectoryType
	}{
		IsAdmin: 1,
		Types:   types,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, templData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error by execute ", err)
	}
}

func (h *EquipmentDirectoryHandler) NewDirectoryPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	name := r.Form.Get("name")
	releaseYear := r.Form.Get("releaseYear")
	eType := r.Form.Get("typeId")
	typeId, _ := strconv.Atoi(eType)

	directory := domain.EquipmentDirectory{
		Name:        name,
		ReleaseYear: releaseYear,
		TypeId:      typeId,
	}

	err = h.EquipmentDirectoryService.NewDirectory(directory)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка добавления директории: ", err)
		return
	}

	http.Redirect(w, r, "/equipmentDirectory", http.StatusSeeOther)
}

func (h *EquipmentDirectoryHandler) DeleteDirectory(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = h.EquipmentDirectoryService.DeleteDirectory(id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка удаления директории: ", err)
		return
	}

	http.Redirect(w, r, "/equipmentDirectory", http.StatusSeeOther)
}
