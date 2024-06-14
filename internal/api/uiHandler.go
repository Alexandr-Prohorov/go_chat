package api

import (
	"net/http"
	"os"
	"path/filepath"
)

type StaticFile struct {
	Path        string
	ContentType string
}

func NewStaticFile(path string, contentType string) *StaticFile {
	return &StaticFile{
		Path:        path,
		ContentType: contentType,
	}
}

func (s *StaticFile) StaticFileHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("ui", s.Path)
	file, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Could not read file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", s.ContentType)
	w.Write(file)
}
