package main

import (
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	http.HandleFunc("/", rootHandler)

	http.ListenAndServe(":8080", nil)
}
