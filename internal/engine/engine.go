package engine

/*
#cgo LDFLAGS: -lOpenCL
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

func New() *Engine {
	return &Engine{eng: C.create_engine()}
}
func (e *Engine) RevertColor(pixels unsafe.Pointer, size int) {
	e.mu.Lock()
	defer e.mu.Unlock()
	C.revert_color(e.eng, (*C.uchar)(pixels), C.int(size))
}
