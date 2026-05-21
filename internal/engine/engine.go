package engine

/*
#include "engine.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

func AnalyseImage(pixels unsafe.Pointer, size int) {
	C.get_img((*C.uchar)(pixels), (C.int)(size))
	runtime.KeepAlive(pixels)
}
