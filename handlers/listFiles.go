package handlers

import (
	// "fmt"
	"net/http"
	"os"
	"encoding/json"
)

const uploadDir = "./uploads"

func ListFilesHandler(w http.ResponseWriter, r *http.Request) {

	// Get all files
	entries, err := os.ReadDir(uploadDir)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Filter directories
	var files []string
	for _, e := range entries {
		if !e.IsDir() {
			files = append(files, e.Name())
		}
	}

	// Set response header
	w.Header().Set("Content-Type", "application/json")

	// Convert file list to JSON and send response
	json.NewEncoder(w).Encode(files)
}
