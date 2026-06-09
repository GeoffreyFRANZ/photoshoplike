package api

/*
#include <stdlib.h>
*/
import "C"
import (
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"photoshop-like/internal/session"
	"unsafe"
)

type Pixel struct {
	R, G, B, A uint8
}

func (s *Server) Upload(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("img")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bounds, pixels, width, height, err := decodeToPixel(file)

	_, err = s.session.Create(r, w, session.PixelsData{
		Pixels: pixels,
		Size:   len(pixels),
		Width:  width,
		Height: height,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	new_image := image.NewRGBA(bounds)

	copy(new_image.Pix, pixels)
	w.Header().Set("Content-Type", "image/jpg")
	err = jpeg.Encode(w, new_image, nil)
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{uint8(r / 257), uint8(g / 257), uint8(b / 257), uint8(a / 257)}
}

func decodeToPixel(file io.Reader) (image.Rectangle, []byte, int, int, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return image.Rectangle{}, nil, 0, 0, err
	}
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	pixels := make([]Pixel, width*height)
	i := 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixels[i] = rgbaToPixel(img.At(x, y).RGBA())
			i++
		}
	}
	pixelBytes := unsafe.Slice((*byte)(unsafe.Pointer(&pixels[0])), len(pixels)*4)

	return bounds, pixelBytes, width, height, err
}
