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
	"github.com/Thomas3246/EquipAccounting/pkg/templateloader"
	"github.com/go-chi/chi/v5"
)

type HardwareHandler struct {
	hardwareService service.HardwareService
}

func NewHardwareHandler(hardwareService *service.HardwareService) *HardwareHandler {
	return &HardwareHandler{hardwareService: *hardwareService}
}

func (h *HardwareHandler) Units(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templateloader.GetTemplatePath("base.html"), templateloader.GetTemplatePath("hardwareList.html"))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Error by parse", err)
		return
	}

	hType := strings.ToLower(r.URL.Query().Get("type"))
	if hType == "" {
		http.Redirect(w, r, "/hardware?type=cpu", http.StatusFound)
		return
	}

	units, err := h.hardwareService.GetUnitsByType(hType)
	if err != nil {
		if err == service.ErrInvalidParameter {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("Ошибка при получении комплектующих: ", err)
			return
		}
	}

	templData := struct {
		IsAdmin     int
		CurrentType string
		Units       []domain.Unit
	}{
		IsAdmin:     1,
		CurrentType: hType,
		Units:       units,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, templData)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Println("Template Execute Error: ", err)
	}
}

func (h *HardwareHandler) UnitGet(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templateloader.GetTemplatePath("base.html"), templateloader.GetTemplatePath("hardware.html"))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Error by parse", err)
		return
	}

	hType := strings.ToLower(chi.URLParam(r, "type"))

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Println(err)
		return
	}

	unit, err := h.hardwareService.GetUnit(hType, id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка получения комплектующего: ", err)
		return
	}

	templData := struct {
		IsAdmin     int
		CurrentType string
		Id          int
		Unit        domain.Unit
	}{
		IsAdmin:     1,
		CurrentType: hType,
		Id:          id,
		Unit:        unit,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, templData)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Println("Template Execute Error: ", err)
	}
}

func (h *HardwareHandler) UnitPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	hType := strings.ToLower(chi.URLParam(r, "type"))

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Println(err)
		return
	}

	name := r.Form.Get("name")

	err = h.hardwareService.UpdateName(hType, id, name)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка обновления имени: ", err)
		return
	}

	url := fmt.Sprintf("/hardware?type=%s", hType)
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func (h *HardwareHandler) NewUnitGet(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templateloader.GetTemplatePath("base.html"), templateloader.GetTemplatePath("newHardware.html"))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Error by parse", err)
		return
	}

	templData := struct {
		IsAdmin int
	}{
		IsAdmin: 1,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, templData)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Println("Template Execute Error: ", err)
	}
}

func (h *HardwareHandler) NewUnitPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	name := r.Form.Get("name")
	hType := r.Form.Get("unit_type")

	err = h.hardwareService.NewUnit(hType, name)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при добавлении комплектующего: ", err)
		return
	}

	http.Redirect(w, r, "/hardware?type=cpu", http.StatusSeeOther)
}

func (h *HardwareHandler) DeleteUnit(w http.ResponseWriter, r *http.Request) {
	hType := chi.URLParam(r, "type")
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = h.hardwareService.DeleteUnit(hType, id)
	if err != nil {
		if err == service.ErrInvalidParameter {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка удаления комплектующего: ", err)
		return
	}

	url := fmt.Sprintf("/hardware?type=%s", hType)
	http.Redirect(w, r, url, http.StatusSeeOther)
}
