package winapi

import (
	"unsafe"
)

// https://msdn.microsoft.com/en-us/library/ms646256.aspx
func DragDetect(hwnd uintptr, pt POINT) (uintptr, error) {
	r1, _, err := procDragDetect.Call(hwnd, uintptr(unsafe.Pointer(&pt)))
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/hh447467.aspx
func EnableMouseInPointer(fEnable uintptr) (uintptr, error) {
	r1, _, err := procEnableMouseInPointer.Call(fEnable)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646257.aspx
func GetCapture() (uintptr, error) {
	r1, _, err := procGetCapture.Call()
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646258.aspx
func GetDoubleClickTime() (uintptr, error) {
	r1, _, err := procGetDoubleClickTime.Call()
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646259.aspx
func GetMouseMovePointsEx(cbSize uintptr, lppt *MOUSEMOVEPOINT, lpptBuf []MOUSEMOVEPOINT, nBufPoints uintptr, resolution uintptr) (uintptr, error) {
	r1, _, err := procGetMouseMovePointsEx.Call(cbSize, uintptr(unsafe.Pointer(lppt)), uintptr(unsafe.Pointer(&lpptBuf[0])), nBufPoints, resolution)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646260.aspx
func Mouse_event(dwFlags uintptr, dx uintptr, dy uintptr, dwData uintptr, dwExtraInfo uintptr) (uintptr, error) {
	r1, _, err := procMouse_event.Call(dwFlags, dx, dy, dwData, dwExtraInfo)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646261.aspx
func ReleaseCapture() (uintptr, error) {
	r1, _, err := procReleaseCapture.Call()
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646262.aspx
func SetCapture(hWnd uintptr) (uintptr, error) {
	r1, _, err := procSetCapture.Call(hWnd)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646263.aspx
func SetDoubleClickTime(uInterval uintptr) (uintptr, error) {
	r1, _, err := procSetDoubleClickTime.Call(uInterval)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646264.aspx
func SwapMouseButton(fSwap uintptr) (uintptr, error) {
	r1, _, err := procSwapMouseButton.Call(fSwap)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646265.aspx
func TrackMouseEvent(lpEventTrack *TRACKMOUSEEVENT) (uintptr, error) {
	r1, _, err := procTrackMouseEvent.Call(uintptr(unsafe.Pointer(lpEventTrack)))
	return r1, err
}
