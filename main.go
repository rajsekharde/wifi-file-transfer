package main

import (
	// "fmt"
	"github.com/fatih/color"
	"github.com/rajsekharde/wifi-file-transfer/handlers"
	"github.com/rajsekharde/wifi-file-transfer/utils"
	"log"
	"net/http"
	"strconv"
	"time"
	"os"
	"path/filepath"
)

const PORT = 8000

var (
	// red   = color.RGB(255, 0, 0).SprintFunc()
	green = color.RGB(0, 255, 0).SprintFunc()
	blue  = color.RGB(100, 150, 255).SprintFunc()
)

func main() {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	realPath, err := filepath.EvalSymlinks(exePath)
	if err != nil {
		log.Fatal(err)
	}

	baseDir := filepath.Dir(realPath)
	staticDir := filepath.Join(baseDir, "static")

	// Serve frontend on "/"
	// fs := http.FileServer(http.Dir("./static"))
	fs := http.FileServer(http.Dir(staticDir))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		fs.ServeHTTP(w, r)
		// Print log
		log.Printf("%s %s %s %s   %s",
			blue(r.Method),
			r.URL.Path,
			green(http.StatusOK),
			green("OK"),
			time.Since(startTime),
		)
	})

	http.HandleFunc("/files", handlers.ListFilesHandler)
	http.HandleFunc("/upload", handlers.UploadHandler)
	http.HandleFunc("/download/", handlers.DownloadHandler)

	// fmt.Printf("Server running on port %d\n", PORT)
	utils.DisplayAsciiLogo(PORT)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), nil))
}
