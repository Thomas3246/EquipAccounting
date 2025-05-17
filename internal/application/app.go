package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Thomas3246/EquipAccounting/internal/application/service"
	"github.com/Thomas3246/EquipAccounting/internal/presentation"
	"github.com/Thomas3246/EquipAccounting/internal/presentation/handler"
)

type App struct {
	db *sql.DB
}

func NewApp(db *sql.DB) *App {
	return &App{db: db}
}

func (a *App) Start() {
	appService := service.NewAppService(a.db)

	appHandler := handler.NewAppHandler(appService)

	router := presentation.NewRouter(appHandler)

	log.Println("Started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
