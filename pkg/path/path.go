package path

import (
	"log"
	"os"
	"path/filepath"
)

func GetStaticPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// Поднимаемся на уровень выше из cmd/
	return filepath.Join(dir, "..", "static")
}
