package handlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	// "strconv"
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

	// Print logs
	if r.Header.Get("Range") == "" {
		info, _ := os.Stat(filePath)

		log.Printf("%s downloaded %s%s%s\n",
			grey(fileName),
			grey("("),
			grey(convertBytes(int(info.Size()))),
			grey(")"),
		)

		log.Printf("%s %s %s %s   %s",
			blue(r.Method),
			r.URL.Path,
			green(http.StatusOK),
			green("OK"),
			time.Since(startTime),
		)
	}
}
