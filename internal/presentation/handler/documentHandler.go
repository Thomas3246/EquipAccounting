package handler

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"fmt"

	"github.com/Thomas3246/EquipAccounting/internal/application/service"
	"github.com/Thomas3246/EquipAccounting/internal/domain"
	"github.com/Thomas3246/EquipAccounting/pkg/templateloader"
	"github.com/go-chi/chi/v5"
)

type DocumentHandler struct {
	DocumentService service.DocumentService
	UserService     service.UserService
}

func NewDocumentHandler(docService *service.DocumentService, userService *service.UserService) *DocumentHandler {
	return &DocumentHandler{
		DocumentService: *docService,
		UserService:     *userService,
	}
}

func (h *DocumentHandler) AddDocumentGet(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templateloader.GetTemplatePath("base.html"), templateloader.GetTemplatePath("addDocument.html"))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Error by parse", err)
		return
	}

	reqIdStr := chi.URLParam(r, "id")
	reqId, err := strconv.Atoi(reqIdStr)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	docTypes, err := h.DocumentService.GetDocumentTypes()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении типов документов: ", err)
		return
	}

	templData := struct {
		IsAdmin       int
		RequestId     int
		DocumentTypes []domain.DocumentType
		DocSizeError  string
	}{
		IsAdmin:       1,
		RequestId:     reqId,
		DocumentTypes: docTypes,
		DocSizeError:  "",
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, templData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error by execute ", err)
	}
}

func (h *DocumentHandler) AddDocumentPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, "Ошибка загрузки файла", http.StatusBadRequest)
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
	userLogin := parts[0]

	requestIdStr := chi.URLParam(r, "id")
	requestId, err := strconv.Atoi(requestIdStr)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	typeIdStr := r.PostFormValue("document_type_id")
	typeId, _ := strconv.Atoi(typeIdStr)

	file, handler, err := r.FormFile("document_file")
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	if handler.Size > 16<<20 {
		tmpl, err := template.ParseFiles(templateloader.GetTemplatePath("base.html"), templateloader.GetTemplatePath("addDocument.html"))
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Println("Error by parse", err)
			return
		}

		reqIdStr := chi.URLParam(r, "id")
		reqId, err := strconv.Atoi(reqIdStr)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		docTypes, err := h.DocumentService.GetDocumentTypes()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("Ошибка при получении типов документов: ", err)
			return
		}

		templData := struct {
			IsAdmin       int
			RequestId     int
			DocumentTypes []domain.DocumentType
			DocSizeError  string
		}{
			IsAdmin:       1,
			RequestId:     reqId,
			DocumentTypes: docTypes,
			DocSizeError:  "Файл слишком большой",
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		err = tmpl.Execute(w, templData)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("Error by execute ", err)
		}
	}

	user, err := h.UserService.GetUserByLogin(userLogin)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении пользователя")
		return
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка чтения файла: ", err)
		return
	}

	document := domain.Document{
		RequestId: requestId,
		Type:      typeId,
		File:      fileBytes,
		UserId:    user.Id,
		Name:      handler.Filename,
	}

	err = h.DocumentService.AddDocument(document)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при добавлении файла в БД: ", err)
		return
	}

	pathToRedirect := "/request/" + requestIdStr
	http.Redirect(w, r, pathToRedirect, http.StatusSeeOther)
}

func (h *DocumentHandler) DownloadDocument(w http.ResponseWriter, r *http.Request) {
	docIdStr := chi.URLParam(r, "id")
	docId, err := strconv.Atoi(docIdStr)
	if err != nil {
		http.Error(w, "BadRequest", http.StatusBadRequest)
		return
	}

	document, err := h.DocumentService.GetDocument(docId)
	if err != nil {
		if err == service.ErrNotFound {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении файла: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=\""+document.Name+"\"")
	w.Header().Set("Content-Transfer-Encoding", "binary")

	w.Write(document.File)
}

func (h *DocumentHandler) DeleteDocument(w http.ResponseWriter, r *http.Request) {
	docIdStr := chi.URLParam(r, "id")
	docId, err := strconv.Atoi(docIdStr)
	if err != nil {
		http.Error(w, "BadRequest", http.StatusBadRequest)
		return
	}

	document, err := h.DocumentService.GetDocument(docId)
	if err != nil {
		if err == service.ErrNotFound {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при получении файла: ", err)
		return
	}

	err = h.DocumentService.DeleteDocument(docId)
	if err != nil {
		if err == service.ErrNotFound {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Ошибка при удалении файла: ", err)
		return
	}

	redirectPath := "/request/" + fmt.Sprint(document.RequestId)
	http.Redirect(w, r, redirectPath, http.StatusSeeOther)
}
