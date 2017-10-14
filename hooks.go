package winapi

import (
	"unsafe"
)

// https://msdn.microsoft.com/en-us/library/ms644973.aspx
func CallMsgFilter(lpMsg *MSG, nCode int32) (uintptr, error) {
	r1, _, err := procCallMsgFilter.Call(uintptr(unsafe.Pointer(lpMsg)), int32ToUintptr(nCode))
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644974.aspx
func CallNextHookEx(hhk uintptr, nCode int32, wParam uintptr, lParam uintptr) (uintptr, error) {
	r1, _, err := procCallNextHookEx.Call(hhk, int32ToUintptr(nCode), wParam, lParam)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644990.aspx
func SetWindowsHookEx(idHook int32, lpfn uintptr, hMod uintptr, dwThreadId uintptr) (uintptr, error) {
	r1, _, err := procSetWindowsHookEx.Call(int32ToUintptr(idHook), lpfn, hMod, dwThreadId)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644993.aspx
func UnhookWindowsHookEx(hhk uintptr) (uintptr, error) {
	r1, _, err := procUnhookWindowsHookEx.Call(hhk)
	return r1, err
}
