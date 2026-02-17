package handlers

import (
	// "fmt"
	// "github.com/fatih/color"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	// "strconv"
	"time"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	startTime := time.Now()

	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 10GB file limit
	r.Body = http.MaxBytesReader(w, r.Body, 10<<30)

	// Use 32MB RAM for multipart parsing
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["files"]

	if len(files) == 0 {
		http.Error(w, "No files uploaded", http.StatusBadRequest)
		return
	}

	for _, header := range files {
		file, err := header.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dst, err := os.Create(filepath.Join(uploadDir, header.Filename))
		if err != nil {
			file.Close()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := io.Copy(dst, file); err != nil {
			file.Close()
			dst.Close()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// log.Printf("%s uploaded (%d bytes)\n", header.Filename, header.Size)
		log.Printf("%s uploaded %s%s%s\n",
			grey(header.Filename),
			grey("("),
			// grey(strconv.Itoa(int(header.Size))),
			grey(convertBytes(int(header.Size))),
			grey(")"),
		)

		file.Close()
		dst.Close()
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("uploaded"))

	// Print log
	log.Printf("%s %s %s %s  %s",
		blue(r.Method),
		r.URL.Path,
		green(http.StatusCreated),
		green("CREATED"),
		time.Since(startTime),
	)
}
