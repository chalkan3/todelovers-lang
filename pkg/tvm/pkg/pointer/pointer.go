package pointer

import "unsafe"

func StringPtr(str string) (unsafe.Pointer, int) {
	return unsafe.Pointer(&str), len(str)
}
