package main

import (
	"log"

	app "github.com/Thomas3246/EquipAccounting/internal/application"
	"github.com/Thomas3246/EquipAccounting/internal/infrastructure/database/sqlite"
)

func main() {

	db, err := sqlite.InitDB()
	if err != nil {
		log.Fatalf("Произошла ошибка подключения к базе данных: %v", err)
	}

	application := app.NewApp(db)
	application.Start()

}
