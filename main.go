package main

import (
	"fmt"
	"net/http"
	"strconv"
)

const PORT = 8000

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r* http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	fmt.Printf("Server running on port %d\n", PORT)
	http.ListenAndServe(":"+strconv.Itoa(PORT), nil)
}