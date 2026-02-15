package handlers

import (
	// "fmt"
	// "github.com/fatih/color"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	startTime := time.Now()

	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Extract file from request
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	defer file.Close()

	// Create destination file
	dst, err := os.Create(filepath.Join(uploadDir, header.Filename))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("uploaded"))

	// Print log
	log.Printf("%s %s %s %s   %s %s   %s",
		blue(r.Method),
		r.URL.Path,
		green(http.StatusCreated),
		green("CREATED"),
		grey(header.Size),
		grey("bytes"),
		time.Since(startTime),
	)
}
