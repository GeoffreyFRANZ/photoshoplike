package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./template/index.html")
	})
	mux.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		Upload(w, r)
	})
	mux.HandleFunc("/reverse", func(w http.ResponseWriter, r *http.Request) {
		reverse_pixels(w, r)
	})

	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
