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
func RevertingColor(pixels unsafe.Pointer, size int) {
	C.revert_color((*C.uchar)(pixels), (C.int)(size))
}
