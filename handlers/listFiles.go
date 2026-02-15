package handlers

import (
	// "fmt"
	"encoding/json"
	"github.com/fatih/color"
	"log"
	"net/http"
	"os"
)

const uploadDir = "./uploads"

var (
	red   = color.RGB(255, 0, 0).SprintFunc()
	green = color.RGB(0, 255, 0).SprintFunc()
	blue  = color.RGB(200, 200, 255).SprintFunc()
)

func ListFilesHandler(w http.ResponseWriter, r *http.Request) {

	// Get all files
	entries, err := os.ReadDir(uploadDir)
	if err != nil {
		log.Printf("%s %s %s ERROR: %v",
		blue(r.Method),
		r.URL.Path,
		red(http.StatusInternalServerError),
		red(err),
	)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Filter directories
	var files []string
	for _, e := range entries {
		if !e.IsDir() {
			files = append(files, e.Name())
		}
	}

	// Print log
	log.Printf("%s %s %s %s",
		blue(r.Method),
		r.URL.Path,
		green(http.StatusOK),
		green("OK"),
	)

	// Set response header
	w.Header().Set("Content-Type", "application/json")

	// Convert file list to JSON and send response
	json.NewEncoder(w).Encode(files)
}
