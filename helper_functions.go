package winapi

import (
	"unsafe"
	"golang.org/x/sys/windows"
)

func int32ToUintptr(v int32) uintptr {
	if v < 0 {
		return uintptr(uint32(1<<32 - 1) - uint32(-v) + 1)
	}
	return uintptr(v)
}

func strToUintptr(str string) (uintptr, error) {
	if str == "" {
		return 0, nil
	}
	utfp, err := windows.UTF16PtrFromString(str)
	if err != nil {
		return 0, err
	}
	return uintptr(unsafe.Pointer(utfp)), nil
}
