package main

import (
	"net/http"
	"sync"

	"github.com/gorilla/sessions"
)

var sessionKey = []byte("sessionKey")
var store = sessions.NewCookieStore(sessionKey)
var m sync.Map

func createSession(r *http.Request, w http.ResponseWriter, pixels []byte, size int, height int, width int) {
	pixelsData := PixelsData{pixels, size, width, height}
	sessionsID, err := store.New(r, "SessionsID")
	if err != nil {
		return
	}
	m.Store(sessionsID.ID, pixelsData)
	err = store.Save(r, w, sessionsID)
	if err != nil {
		return
	}
}

type PixelsData struct {
	Pixels []byte
	Size   int
	Width  int
	Height int
}
