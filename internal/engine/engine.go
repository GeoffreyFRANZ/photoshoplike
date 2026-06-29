package engine

/*
#cgo LDFLAGS: -lOpenCL -lopenvino_c
#include "engine.h"
*/
import "C"
import (
	"sync"
	"unsafe"
)

type Engine struct {
	eng *C.opencl_engine
	mu  sync.Mutex
}
type OpenVino struct {
	vino *C.openvino_engine
	mu   sync.Mutex
}

func NewEngine() *Engine {
	return &Engine{eng: C.create_engine()}
}

func NewOpenVino() *OpenVino {
	return &OpenVino{vino: C.create_openvino_engine()}
}
func (e *Engine) RevertColor(pixels unsafe.Pointer, size int) {
	e.mu.Lock()
	defer e.mu.Unlock()
	C.revert_color(e.eng, (*C.uchar)(pixels), C.int(size))
}

func (o *OpenVino) ContrastColor(pointer unsafe.Pointer, size, height, width int) {
	o.mu.Lock()
	defer o.mu.Unlock()
	C.contrast_color(o.vino, (*C.uchar)(pointer), C.int(size), C.int(height), C.int(width))
}
