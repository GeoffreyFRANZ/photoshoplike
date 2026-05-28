package main

import (
	"fmt"
	"net/http"
)

func reversing_pixels(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "SessionsID")
	if err != nil {
		return
	}
	data, ok := m.Load(session.ID)
	if !ok {
		return
	}
	pixels, _ := m.Load(session.ID)
	fmt.Println(pixels)

	w.Header().Set("Content-Type", "image/jpg")
	w.Write(data.(PixelsData).Pixels)
}
