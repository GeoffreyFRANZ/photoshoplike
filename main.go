package main

import (
	"log"
	"net/http"
	"photoshop-like/internal/engine"
)

func main() {
	mux := http.NewServeMux()
	server := Server{engine.New()}
	if server.engine == nil {
		panic("failed to create engine")
	}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./template/index.html")
	})
	mux.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		Upload(w, r)
	})
	mux.HandleFunc("/reverse-color", server.reverseColorHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}

type Server struct {
	engine *engine.Engine
}
