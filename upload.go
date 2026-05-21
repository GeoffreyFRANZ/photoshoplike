package main

/*
#cgo LDFLAGS: -lOpenCL
*/
import "C"
import (
	"image"
	"log"
	"net/http"
	"sentinel-go/internal/engine"
	"unsafe"
)

import (
	_ "image/jpeg"
	_ "image/png"
)

type Pixel struct {
	R, G, B, A uint8
}

func Upload(w http.ResponseWriter, r *http.Request) int {
	file, fileheader, err := r.FormFile("img")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return http.StatusBadRequest
	}
	log.Println(file)
	log.Println(fileheader.Filename)
	log.Println(fileheader.Header)
	log.Println("size : ", fileheader.Size)
	img, _, err := image.Decode(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return http.StatusBadRequest
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
	cData := C.CBytes(pixelBytes)
	defer C.free(cData)
	engine.AnalyseImage(cData, len(pixels)*4)
	return http.StatusOK
}
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{uint8(r / 257), uint8(g / 257), uint8(b / 257), uint8(a / 257)}
}
