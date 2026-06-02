package engine

/*
#cgo LDFLAGS: -lOpenCL
#include "engine.h"
*/
import "C"
import "unsafe"

type Engine struct {
	eng *C.opencl_engine
}

func New() *Engine {
	return &Engine{C.create_engine()}
}
func (e *Engine) RevertColor(pixels unsafe.Pointer, size int) {
	C.revert_color(e.eng, (*C.uchar)(pixels), C.int(size))
}
