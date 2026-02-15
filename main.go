package main

import (
	"fmt"
	"net/http"
	"strconv"
)

const PORT = 8000

func main() {
	frontendHandler := http.FileServer(http.Dir("./static"))

	http.Handle("/", frontendHandler)

	fmt.Printf("Server running on port %d\n", PORT)
	http.ListenAndServe(":"+strconv.Itoa(PORT), nil)
}