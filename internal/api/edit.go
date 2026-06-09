package api

import (
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"unsafe"
)

func (s *Server) ReverseColorHandler(w http.ResponseWriter, r *http.Request) {
	_, err := io.Copy(io.Discard, r.Body)
	if err != nil {
		return
	}
	sess, b := s.session.Load(r)
	if b == false {
		return
	}

	pixels := sess.Pixels
	size := sess.Size
	width := sess.Width
	height := sess.Height
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	s.engine.RevertColor(unsafe.Pointer(&pixels[0]), size)
	copy(img.Pix, pixels)
	w.Header().Set("Content-Type", "image/jpg")
	err = jpeg.Encode(w, img, nil)
}
