package handlers

import (
	// "fmt"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type FileInfo struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
}

func ListFilesHandler(w http.ResponseWriter, r *http.Request) {

	startTime := time.Now()

	// Get all files
	entries, err := os.ReadDir(uploadDir)
	if err != nil {
		log.Printf("%s %s %s ERROR: %v   %s",
			blue(r.Method),
			r.URL.Path,
			red(http.StatusInternalServerError),
			red(err),
			time.Since(startTime),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Filter directories
	var files []FileInfo
	for _, entry := range entries {
		if !entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				continue
			}
			files = append(files, FileInfo{
				Name: entry.Name(),
				Size: info.Size(),
			})
		}
	}

	// Set response header
	w.Header().Set("Content-Type", "application/json")

	// Convert file list to JSON and send response
	json.NewEncoder(w).Encode(files)

	// Print log
	log.Printf("%s %s %s %s   %s",
		blue(r.Method),
		r.URL.Path,
		green(http.StatusOK),
		green("OK"),
		time.Since(startTime),
	)
}
