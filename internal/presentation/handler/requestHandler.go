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

type RequestHandler struct {
	reqService  service.RequestService
	userService service.UserService
	eqService   service.EquipmentService
	docService  service.DocumentService
}

func NewRequestHandler(
	reqService *service.RequestService,
	userService *service.UserService,
	eqService *service.EquipmentService,
	docServise *service.DocumentService,
) *RequestHandler {
	return &RequestHandler{
		reqService:  *reqService,
		userService: *userService,
		eqService:   *eqService,
		docService:  *docServise,
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

	isAdmin, err := strconv.Atoi(parts[1])
	if err != nil {
		http.Error(w, "Invalid Cookie Value", http.StatusInternalServerError)
		log.Println("Ошибка чтения значения isAdmin в cookie: ", err)
		return
	}

	templData := struct {
		Requests  []domain.RequestView
		UserLogin string
		Flag      string
		IsAdmin   int
	}{
		Requests:  requests,
		UserLogin: parts[0],
		Flag:      "allactive",
		IsAdmin:   isAdmin,
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

	isAdmin, err := strconv.Atoi(parts[1])
	if err != nil {
		http.Error(w, "Invalid Cookie Value", http.StatusInternalServerError)
		log.Println("Ошибка чтения значения isAdmin в cookie: ", err)
		return
	}

	templData := struct {
		Requests  []domain.RequestView
		UserLogin string
		Flag      string
		IsAdmin   int
	}{
		Requests:  requests,
		UserLogin: parts[0],
		Flag:      "alluseractive",
		IsAdmin:   isAdmin,
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

	isAdmin, err := strconv.Atoi(parts[1])
	if err != nil {
		http.Error(w, "Invalid Cookie Value", http.StatusInternalServerError)
		log.Println("Ошибка чтения значения isAdmin в cookie: ", err)
		return
	}

	requests, err := h.reqService.GetAllClosed(cookie.Value)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при извлечении заявок: ", err)
		return
	}

	templData := struct {
		Requests  []domain.RequestView
		UserLogin string
		Flag      string
		IsAdmin   int
	}{
		Requests:  requests,
		UserLogin: parts[0],
		Flag:      "allclosed",
		IsAdmin:   isAdmin,
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

	isAdmin, err := strconv.Atoi(parts[1])
	if err != nil {
		http.Error(w, "Invalid Cookie Value", http.StatusInternalServerError)
		log.Println("Ошибка чтения значения isAdmin в cookie: ", err)
		return
	}

	templData := struct {
		Requests  []domain.RequestView
		UserLogin string
		Flag      string
		IsAdmin   int
	}{
		Requests:  requests,
		UserLogin: parts[0],
		Flag:      "alluserclosed",
		IsAdmin:   isAdmin,
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

	isAdmin, err := strconv.Atoi(parts[1])
	if err != nil {
		http.Error(w, "Invalid Cookie Value", http.StatusInternalServerError)
		log.Println("Ошибка чтения значения isAdmin в cookie: ", err)
		return
	}

	requestTypes, err := h.reqService.GetRequestTypes()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении типов заявок: ", err)
		return
	}

	equipment, err := h.eqService.GetAvailableEquipment(cookie.Value)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении доступного оборудования: ", err)
		return
	}

	templData := struct {
		UserLogin    string
		RequestTypes []domain.RequestType
		Equipment    []domain.EquipmentView
		IsAdmin      int
	}{
		UserLogin:    parts[0],
		RequestTypes: requestTypes,
		Equipment:    equipment,
		IsAdmin:      isAdmin,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, templData)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Println("Template Execute Error: ", err)
	}
}

func (h *RequestHandler) NewRequestPost(w http.ResponseWriter, r *http.Request) {

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

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Println("Ошибка при парсинге формы: ", err)
		return
	}

	requestTypeId, err := strconv.Atoi(r.FormValue("request_type_id"))
	if err != nil {
		http.Error(w, "Invalid Request Type Id", http.StatusBadRequest)
		return
	}

	equipmentId, err := strconv.Atoi(r.FormValue("equipment_id"))
	if err != nil {
		http.Error(w, "Invalid Equipment Id", http.StatusBadRequest)
		return
	}

	description := r.FormValue("description")

	user, err := h.userService.GetUserByLogin(parts[0])
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = h.reqService.NewRequest(requestTypeId, description, user.Id, equipmentId)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при добавлении заявки: ", err)
		return
	}

	err = h.eqService.ChangeEquipmentStatus(equipmentId)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при изменении статуса оборудования: ", err)
		return
	}

	http.Redirect(w, r, "/allactive", http.StatusSeeOther)
}

func (h *RequestHandler) RequestEditGet(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templateloader.GetTemplatePath("base.html"), templateloader.GetTemplatePath("request.html"))
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

	requestId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	files, err := h.docService.GetDocumentsViewForRequest(requestId)
	if err != nil {
		http.Error(w, "Internal Server Errror", http.StatusInternalServerError)
		log.Println("Ошибка при получении документов на заявку: ", err)
		return
	}

	request, err := h.reqService.GetRequestById(requestId)
	if err != nil {
		if err == service.ErrNotFound {
			http.Error(w, "Request Not Found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении запроса: ", err)
		return
	}

	isAdmin, err := strconv.Atoi(parts[1])
	if err != nil {
		http.Error(w, "Invalid Cookie Value", http.StatusInternalServerError)
		log.Println("Ошибка чтения значения isAdmin в cookie: ", err)
		return
	}

	equipment, err := h.eqService.GetAvailableEquipment(cookie.Value)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении оборудования: ", err)
		return
	}

	allEquipment, err := h.eqService.GetAllEquipment()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении оборудования: ", err)
		return
	}

	requestTypes, err := h.reqService.GetRequestTypes()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении типов заявок: ", err)
		return
	}

	user, err := h.userService.GetUserByLogin(parts[0])
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении автора заявки: ", err)
		return
	}

	results, err := h.reqService.GetRequestResults()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении результатов заявки: ", err)
		return
	}

	templData := struct {
		Request      domain.Request
		IsAdmin      int
		Equipment    []domain.EquipmentView
		Types        []domain.RequestType
		Author       domain.User
		Documents    []domain.DocumentView
		AllEquipment []domain.EquipmentView
		Results      []domain.RequestResult
	}{
		Request:      request,
		IsAdmin:      isAdmin,
		Equipment:    equipment,
		Types:        requestTypes,
		Author:       *user,
		Documents:    files,
		AllEquipment: allEquipment,
		Results:      results,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, templData)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Println("Template Execute Error: ", err)
	}
}

func (h *RequestHandler) RequestEditPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
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

	requestId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	if parts[1] == "0" {
		descr := r.Form.Get("description")
		err := h.reqService.EditDescription(requestId, descr)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("Ошибка изменения описания заявки: ", err)
			return
		}
	}

	if parts[1] == "1" {
		rTypeStr := r.Form.Get("type")
		rType, _ := strconv.Atoi(rTypeStr)
		descr := r.Form.Get("description")
		equipStr := r.Form.Get("equipment")
		equip, _ := strconv.Atoi(equipStr)

		request := domain.Request{
			Id:          requestId,
			Type:        rType,
			Description: descr,
			Equipment:   equip,
		}

		err := h.reqService.EditRequest(request)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("Ошибка при редактировании заявки: ", err)
			return
		}
	}

	http.Redirect(w, r, "/allactive", http.StatusSeeOther)
}

func (h *RequestHandler) CloseRequestGet(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templateloader.GetTemplatePath("base.html"), templateloader.GetTemplatePath("requestClose.html"))
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Println("Template Parse Error: ", err)
		return
	}

	requestId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	results, err := h.reqService.GetRequestResults()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении результатов: ", err)
		return
	}

	templData := struct {
		IsAdmin        int
		RequestId      int
		RequestResults []domain.RequestResult
	}{
		IsAdmin:        1,
		RequestId:      requestId,
		RequestResults: results,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, templData)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Println("Template Execute Error: ", err)
	}
}

func (h *RequestHandler) CloseRequestPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	requestId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	result := r.Form.Get("result_id")
	resultId, _ := strconv.Atoi(result)

	resultDescr := r.Form.Get("resultDescr")
	fmt.Println(resultDescr)

	equipId, err := h.reqService.RequestIsTheOnlyOne(requestId)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка получения заявок оборудования: ", err)
		return
	}
	if equipId > 0 {
		if resultId == 2 {
			err = h.eqService.DecomEquipment(equipId)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				log.Println("Ошибка списания оборудования: ", err)
				return
			}
		} else {
			err = h.eqService.ChangeEquipStatusByResult(equipId, resultId)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				log.Println("Ошибка изменения статуса оборудования: ", err)
				return
			}
		}
	}

	err = h.reqService.CloseRequest(requestId, resultId, resultDescr)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при закрытии заявки: ", err)
		return
	}

	http.Redirect(w, r, "/allactive", http.StatusSeeOther)
}

func (h *RequestHandler) FormReport(w http.ResponseWriter, r *http.Request) {
	requestId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
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
	adminLogin := parts[0]

	report, err := h.reqService.FormReportForRequest(requestId, adminLogin)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(report)
}
