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
		status := Upload(w, r)
		w.WriteHeader(status)
		http.Redirect(w, r, "/", status)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
