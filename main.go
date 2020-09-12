package main

import (
	"net/http"
	"os"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	for _, pair := range os.Environ() {
		w.Write([]byte(pair))
		w.Write([]byte("\n"))
	}
}

func main() {
	http.HandleFunc("/", rootHandler)

	http.ListenAndServe(":8080", nil)
}
