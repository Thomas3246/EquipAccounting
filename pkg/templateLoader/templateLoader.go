package templateloader

import (
	"path/filepath"
	"runtime"
)

func GetTemplatePath(filename string) string {
	_, currentFile, _, _ := runtime.Caller(0)
	projectRoot := filepath.Dir(filepath.Dir(filepath.Dir(currentFile)))
	return filepath.Join(projectRoot, "templates", filename)
}
