package handlers

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/florianwoelki/insta-clone/service.image-storage/files"
	"github.com/gorilla/mux"
)

// Files struct represents a Files handler
type Files struct {
	logger  *log.Logger
	storage files.Storage
}

// NewFiles creates a new Files handler
func NewFiles(storage files.Storage, logger *log.Logger) *Files {
	return &Files{logger: logger, storage: storage}
}

// UploadRest implements the http.Handler interface
func (f *Files) UploadRest(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	filename := vars["filename"]

	f.logger.Println("Handle POST", id, "filename", filename)

	if id == "" || filename == "" {
		f.invalidURI(r.URL.String(), rw)
		return
	}

	f.saveFile(id, filename, rw, r)
}

func (f *Files) UploadMultipart(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(128 * 1024)
	if err != nil {
		f.logger.Printf("Something went wrong while parsing multipart form %v", err)
		http.Error(rw, "Expected multipart form data", http.StatusBadRequest)
		return
	}

	id := r.FormValue("id")
	f.logger.Println("Process form for id", id)
}

func (f *Files) invalidURI(uri string, rw http.ResponseWriter) {
	f.logger.Fatal("Invalid path", uri)
	http.Error(rw, "Invalid file path should be in the format: /[id]/[filepath]", http.StatusBadRequest)
}

func (f *Files) saveFile(id, path string, rw http.ResponseWriter, r *http.Request) {
	f.logger.Println("Save file for product", "id", id, "path", path)

	// try to save file
	filepath := filepath.Join(id, path)
	err := f.storage.Save(filepath, r.Body)
	if err != nil {
		f.logger.Fatal("Unable to save file:", err)
		http.Error(rw, "Unable to save file", http.StatusInternalServerError)
	}
}
