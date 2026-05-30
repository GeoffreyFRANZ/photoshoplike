package main

import (
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"photoshop-like/internal/engine"
	"unsafe"
)

func reversing_pixels(w http.ResponseWriter, r *http.Request) {
	io.Copy(io.Discard, r.Body)
	session, err := store.Get(r, "SessionsID")
	if err != nil {
		return
	}
	data, ok := m.Load(session.ID)
	if !ok {
		return
	}
	pixels := data.(PixelsData).Pixels
	size := data.(PixelsData).Size
	width := data.(PixelsData).Width
	height := data.(PixelsData).Height
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	engine.RevertingColor(unsafe.Pointer(&pixels[0]), size)
	copy(img.Pix, pixels)
	w.Header().Set("Content-Type", "image/jpg")
	err = jpeg.Encode(w, img, nil)
}
