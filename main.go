package main

import (
	"log"
	"net/http"
	"os"
	"photoshop-like/internal/api"
	"photoshop-like/internal/engine"
	"photoshop-like/internal/session"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	eng := engine.NewEngine()
	vino := engine.NewOpenVino()
	sess := session.New([]byte(os.Getenv("SESSION_KEY")))
	if eng == nil {
		log.Fatal("failed to create engine")
	}
	server := api.New(eng, vino, sess)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./template/index.html")
	})
	mux.HandleFunc("/upload", server.Upload)
	mux.HandleFunc("/reverse-color", server.ReverseColorHandler)
	mux.HandleFunc("/contrast", server.ContrastHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
