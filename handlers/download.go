package handlers

import (
	"net/http"
	"path/filepath"
	"log"
	"os"
	"time"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {

	startTime := time.Now()

	fileName := filepath.Base(r.URL.Path)
	filePath := filepath.Join(uploadDir, fileName)

	// Check if file exists
	if _, err := os.Stat(filePath); err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		log.Printf("%s %s %s   %s",
			blue(r.Method),
			r.URL.Path,
			red(http.StatusNotFound),
			time.Since(startTime),
		)
		return
	}

	// Force browser to download the file
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	http.ServeFile(w, r, filePath)

	info, _ := os.Stat(filePath)
	// Print log
	log.Printf("%s %s %s %s   %s %s   %s",
		blue(r.Method),
		r.URL.Path,
		green(http.StatusOK),
		green("OK"),
		grey(info.Size()),
		grey("bytes"),
		time.Since(startTime),
	)
}
