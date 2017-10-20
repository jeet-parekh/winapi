package winapi

import (
	"unsafe"
)

// https://docs.microsoft.com/en-us/windows/console/getconsolemode
func GetConsoleMode(hConsoleHandle uintptr, lpMode *uintptr) (uintptr, error) {
	r1, _, err := procGetConsoleMode.Call(hConsoleHandle, uintptr(unsafe.Pointer(lpMode)))
	return r1, err
}

// https://docs.microsoft.com/en-us/windows/console/getstdhandle
func GetStdHandle(nStdHandle uintptr) (uintptr, error) {
	r1, _, err := procGetStdHandle.Call(nStdHandle)
	return r1, err
}

// https://docs.microsoft.com/en-us/windows/console/peekconsoleinput
func PeekConsoleInputKey(hConsoleInput uintptr, lpBuffer []INPUT_RECORD_KEY, nLength uintptr, lpNumberOfEventsRead *uintptr) (uintptr, error) {
	r1, _, err := procPeekConsoleInput.Call(hConsoleInput, uintptr(unsafe.Pointer(&lpBuffer[0])), nLength, uintptr(unsafe.Pointer(lpNumberOfEventsRead)))
	return r1, err
}

// https://docs.microsoft.com/en-us/windows/console/peekconsoleinput
func PeekConsoleInputMouse(hConsoleInput uintptr, lpBuffer []INPUT_RECORD_MOUSE, nLength uintptr, lpNumberOfEventsRead *uintptr) (uintptr, error) {
	r1, _, err := procPeekConsoleInput.Call(hConsoleInput, uintptr(unsafe.Pointer(&lpBuffer[0])), nLength, uintptr(unsafe.Pointer(lpNumberOfEventsRead)))
	return r1, err
}

// https://docs.microsoft.com/en-us/windows/console/peekconsoleinput
func PeekConsoleInputWindowBufferSize(hConsoleInput uintptr, lpBuffer []INPUT_RECORD_WINDOW_BUFFER_SIZE, nLength uintptr, lpNumberOfEventsRead *uintptr) (uintptr, error) {
	r1, _, err := procPeekConsoleInput.Call(hConsoleInput, uintptr(unsafe.Pointer(&lpBuffer[0])), nLength, uintptr(unsafe.Pointer(lpNumberOfEventsRead)))
	return r1, err
}

// https://docs.microsoft.com/en-us/windows/console/peekconsoleinput
func PeekConsoleInputMenu(hConsoleInput uintptr, lpBuffer []INPUT_RECORD_MENU, nLength uintptr, lpNumberOfEventsRead *uintptr) (uintptr, error) {
	r1, _, err := procPeekConsoleInput.Call(hConsoleInput, uintptr(unsafe.Pointer(&lpBuffer[0])), nLength, uintptr(unsafe.Pointer(lpNumberOfEventsRead)))
	return r1, err
}

// https://docs.microsoft.com/en-us/windows/console/peekconsoleinput
func PeekConsoleInputFocus(hConsoleInput uintptr, lpBuffer []INPUT_RECORD_FOCUS, nLength uintptr, lpNumberOfEventsRead *uintptr) (uintptr, error) {
	r1, _, err := procPeekConsoleInput.Call(hConsoleInput, uintptr(unsafe.Pointer(&lpBuffer[0])), nLength, uintptr(unsafe.Pointer(lpNumberOfEventsRead)))
	return r1, err
}

// https://docs.microsoft.com/en-us/windows/console/readconsoleinput
func ReadConsoleInputKey(hConsoleInput uintptr, lpBuffer []INPUT_RECORD_KEY, nLength uintptr, lpNumberOfEventsRead *uintptr) (uintptr, error) {
	r1, _, err := procReadConsoleInput.Call(hConsoleInput, uintptr(unsafe.Pointer(&lpBuffer[0])), nLength, uintptr(unsafe.Pointer(lpNumberOfEventsRead)))
	return r1, err
}

// https://docs.microsoft.com/en-us/windows/console/readconsoleinput
func ReadConsoleInputMouse(hConsoleInput uintptr, lpBuffer []INPUT_RECORD_MOUSE, nLength uintptr, lpNumberOfEventsRead *uintptr) (uintptr, error) {
	r1, _, err := procReadConsoleInput.Call(hConsoleInput, uintptr(unsafe.Pointer(&lpBuffer[0])), nLength, uintptr(unsafe.Pointer(lpNumberOfEventsRead)))
	return r1, err
}

// https://docs.microsoft.com/en-us/windows/console/readconsoleinput
func ReadConsoleInputWindowBufferSize(hConsoleInput uintptr, lpBuffer []INPUT_RECORD_WINDOW_BUFFER_SIZE, nLength uintptr, lpNumberOfEventsRead *uintptr) (uintptr, error) {
	r1, _, err := procReadConsoleInput.Call(hConsoleInput, uintptr(unsafe.Pointer(&lpBuffer[0])), nLength, uintptr(unsafe.Pointer(lpNumberOfEventsRead)))
	return r1, err
}

// https://docs.microsoft.com/en-us/windows/console/readconsoleinput
func ReadConsoleInputMenu(hConsoleInput uintptr, lpBuffer []INPUT_RECORD_MENU, nLength uintptr, lpNumberOfEventsRead *uintptr) (uintptr, error) {
	r1, _, err := procReadConsoleInput.Call(hConsoleInput, uintptr(unsafe.Pointer(&lpBuffer[0])), nLength, uintptr(unsafe.Pointer(lpNumberOfEventsRead)))
	return r1, err
}

// https://docs.microsoft.com/en-us/windows/console/readconsoleinput
func ReadConsoleInputFocus(hConsoleInput uintptr, lpBuffer []INPUT_RECORD_FOCUS, nLength uintptr, lpNumberOfEventsRead *uintptr) (uintptr, error) {
	r1, _, err := procReadConsoleInput.Call(hConsoleInput, uintptr(unsafe.Pointer(&lpBuffer[0])), nLength, uintptr(unsafe.Pointer(lpNumberOfEventsRead)))
	return r1, err
}

// https://docs.microsoft.com/en-us/windows/console/setconsolemode
func SetConsoleMode(hConsoleHandle uintptr, dwMode uintptr) (uintptr, error) {
	r1, _, err := procSetConsoleMode.Call(hConsoleHandle, dwMode)
	return r1, err
}
