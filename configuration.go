package winapi

import (
	"unsafe"
)

// https://msdn.microsoft.com/en-us/library/ms724385.aspx
func GetSystemMetrics(nIndex uintptr) (uintptr, error) {
	r1, _, err := procGetSystemMetrics.Call(nIndex)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms724947.aspx
func SystemParametersInfo(uiAction uintptr, uiParam uintptr, pvParam *uintptr, fWinIni uintptr) (uintptr, error) {
	r1, _, err := procSystemParametersInfo.Call(uiAction, uiParam, uintptr(unsafe.Pointer(pvParam)), fWinIni)
	return r1, err
}
