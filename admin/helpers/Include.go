package helpers

import (
	"path/filepath"
)

func Include(path string) []string {
	files, _ := filepath.Glob("admin/views/templates/*.html")
	pathFiles, _ := filepath.Glob("admin/views/" + path + "/*.html")
	for _, file := range pathFiles {
		files = append(files, file)
	}

	return files
}
