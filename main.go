package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/rajsekharde/wifi-file-transfer/handlers"
	"log"
	"net/http"
	"strconv"
	"time"
)

const PORT = 8000

var (
	// red   = color.RGB(255, 0, 0).SprintFunc()
	green = color.RGB(0, 255, 0).SprintFunc()
	blue  = color.RGB(200, 200, 255).SprintFunc()
)

func main() {

	// Serve frontend on "/"
	fs := http.FileServer(http.Dir("./static"))
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

	fmt.Printf("Server running on port %d\n", PORT)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), nil))
}
