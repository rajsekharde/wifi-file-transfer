package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/rajsekharde/wifi-file-transfer/handlers"
)

const PORT = 8000

func main() {
	frontendHandler := http.FileServer(http.Dir("./static"))

	http.Handle("/", frontendHandler)

	http.HandleFunc("/files", handlers.ListFilesHandler)

	fmt.Printf("Server running on port %d\n", PORT)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), nil))
}