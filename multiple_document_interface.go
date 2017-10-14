package winapi

import (
	"unsafe"
)

// https://msdn.microsoft.com/en-us/library/ms644923.aspx
func CreateMDIWindow(lpClassName string, lpWindowName string, dwStyle uintptr, X int32, Y int32, nWidth uintptr, nHeight uintptr, hWndParent uintptr, hInstance uintptr, lParam uintptr) (uintptr, error) {
	_lpClassName, err := stringToUintptr(lpClassName)
	if err != nil {
		return 0, err
	}
	_lpWindowName, err := stringToUintptr(lpWindowName)
	if err != nil {
		return 0, err
	}
	r1, _, err := procCreateMDIWindow.Call(_lpClassName, _lpWindowName, dwStyle, int32ToUintptr(X), int32ToUintptr(Y), nWidth, nHeight, hWndParent, hInstance, lParam)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644924.aspx
func DefFrameProc(hWnd uintptr, hWndMDIClient uintptr, uMsg uintptr, wParam uintptr, lParam uintptr) (uintptr, error) {
	r1, _, err := procDefFrameProc.Call(hWnd, hWndMDIClient, uMsg, wParam, lParam)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644925.aspx
func DefMDIChildProc(hWnd uintptr, uMsg uintptr, wParam uintptr, lParam uintptr) (uintptr, error) {
	r1, _, err := procDefMDIChildProc.Call(hWnd, uMsg, wParam, lParam)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644926.aspx
func TranslateMDISysAccel(hWndClient uintptr, lpMsg *MSG) (uintptr, error) {
	r1, _, err := procTranslateMDISysAccel.Call(hWndClient, uintptr(unsafe.Pointer(lpMsg)))
	return r1, err
}
