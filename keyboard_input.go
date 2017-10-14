package winapi

import (
	"unsafe"
)

// https://msdn.microsoft.com/en-us/library/ms646289.aspx
func ActivateKeyboardLayout(hkl uintptr, Flags uintptr) (uintptr, error) {
	r1, _, err := procActivateKeyboardLayout.Call(hkl, Flags)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646290.aspx
func BlockInput(fBlockIt uintptr) (uintptr, error) {
	r1, _, err := procBlockInput.Call(fBlockIt)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646291.aspx
func EnableWindow(hWnd uintptr, bEnable uintptr) (uintptr, error) {
	r1, _, err := procEnableWindow.Call(hWnd, bEnable)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646292.aspx
func GetActiveWindow() (uintptr, error) {
	r1, _, err := procGetActiveWindow.Call()
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646293.aspx
func GetAsyncKeyState(vKey uintptr) (uintptr, error) {
	r1, _, err := procGetAsyncKeyState.Call(vKey)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646294.aspx
func GetFocus() (uintptr, error) {
	r1, _, err := procGetFocus.Call()
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646295.aspx
func GetKBCodePage() (uintptr, error) {
	r1, _, err := procGetKBCodePage.Call()
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646296.aspx
func GetKeyboardLayout(idThread uintptr) (uintptr, error) {
	r1, _, err := procGetKeyboardLayout.Call(idThread)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646297.aspx
func GetKeyboardLayoutList(nBuff uintptr, lpList []uintptr) (uintptr, error) {
	r1, _, err := procGetKeyboardLayoutList.Call(nBuff, uintptr(unsafe.Pointer(&lpList[0])))
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646298.aspx
func GetKeyboardLayoutName(pwszKLID []uint16) (uintptr, error) {
	r1, _, err := procGetKeyboardLayoutName.Call(uintptr(unsafe.Pointer(&pwszKLID[0])))
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646299.aspx
func GetKeyboardState(lpKeyState []uint8) (uintptr, error) {
	r1, _, err := procGetKeyboardState.Call(uintptr(unsafe.Pointer(&lpKeyState[0])))
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms724336.aspx
func GetKeyboardType(nTypeFlag uintptr) (uintptr, error) {
	r1, _, err := procGetKeyboardType.Call(nTypeFlag)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646300.aspx
func GetKeyNameText(lParam int32, lpString []uint16, cchSize uintptr) (uintptr, error) {
	r1, _, err := procGetKeyNameText.Call(int32ToUintptr(lParam), uintptr(unsafe.Pointer(&lpString[0])), cchSize)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646301.aspx
func GetKeyState(nVirtKey uintptr) (uintptr, error) {
	r1, _, err := procGetKeyState.Call(nVirtKey)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646302.aspx
func GetLastInputInfo(plii *LASTINPUTINFO) (uintptr, error) {
	r1, _, err := procGetLastInputInfo.Call(uintptr(unsafe.Pointer(plii)))
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646303.aspx
func IsWindowEnabled(hWnd uintptr) (uintptr, error) {
	r1, _, err := procIsWindowEnabled.Call(hWnd)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646304.aspx
func Keybd_event(bVk uintptr, bScan uintptr, dwFlags uintptr, dwExtraInfo uintptr) (uintptr, error) {
	r1, _, err := procKeybd_event.Call(bVk, bScan, dwFlags, dwExtraInfo)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646305.aspx
func LoadKeyboardLayout(pwszKLID string, Flags uintptr) (uintptr, error) {
	_pwszKLID, err := stringToUintptr(pwszKLID)
	if err != nil {
		return 0, err
	}
	r1, _, err := procLoadKeyboardLayout.Call(uintptr(unsafe.Pointer(_pwszKLID)), Flags)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646306.aspx
func MapVirtualKey(uCode uintptr, uMapType uintptr) (uintptr, error) {
	r1, _, err := procMapVirtualKey.Call(uCode, uMapType)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646307.aspx
func MapVirtualKeyEx(uCode uintptr, uMapType uintptr, dwhkl *uintptr) (uintptr, error) {
	r1, _, err := procMapVirtualKeyEx.Call(uCode, uMapType, uintptr(unsafe.Pointer(dwhkl)))
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646308.aspx
func OemKeyScan(wOemChar uintptr) (uintptr, error) {
	r1, _, err := procOemKeyScan.Call(wOemChar)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646309.aspx
func RegisterHotKey(hWnd uintptr, id uintptr, fsModifiers uintptr, vk uintptr) (uintptr, error) {
	r1, _, err := procRegisterHotKey.Call(hWnd, id, fsModifiers, vk)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646310.aspx
func SendInputM(nInputs uintptr, pInputs []INPUT_M, cbSize uintptr) (uintptr, error) {
	return sendInput(_I_M, nInputs, pInputs, cbSize)
}

// https://msdn.microsoft.com/en-us/library/ms646310.aspx
func SendInputK(nInputs uintptr, pInputs []INPUT_K, cbSize uintptr) (uintptr, error) {
	return sendInput(_I_K, nInputs, pInputs, cbSize)
}

// https://msdn.microsoft.com/en-us/library/ms646310.aspx
func SendInputHW(nInputs uintptr, pInputs []INPUT_HW, cbSize uintptr) (uintptr, error) {
	return sendInput(_I_HW, nInputs, pInputs, cbSize)
}

// https://msdn.microsoft.com/en-us/library/ms646311.aspx
func SetActiveWindow(hWnd uintptr) (uintptr, error) {
	r1, _, err := procSetActiveWindow.Call(hWnd)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646312.aspx
func SetFocus(hWnd uintptr) (uintptr, error) {
	r1, _, err := procSetFocus.Call(hWnd)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646314.aspx
func SetKeyboardState(lpKeyState []uint8) (uintptr, error) {
	r1, _, err := procSetKeyboardState.Call(uintptr(unsafe.Pointer(&lpKeyState[0])))
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646316.aspx
func ToAscii(uVirtKey uintptr, uScanCode uintptr, lpKeyState []uint8, lpChar []uint16, uFlags uintptr) (uintptr, error) {
	var _lpKeyState uintptr
	if lpKeyState == nil {
		_lpKeyState = 0
	} else {
		_lpKeyState = uintptr(unsafe.Pointer(&lpKeyState[0]))
	}
	r1, _, err := procToAscii.Call(uVirtKey, uScanCode, _lpKeyState, uintptr(unsafe.Pointer(&lpChar[0])), uFlags)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646318.aspx
func ToAsciiEx(uVirtKey uintptr, uScanCode uintptr, lpKeyState []uint8, lpChar []uint16, uFlags uintptr, dwhkl uintptr) (uintptr, error) {
	var _lpKeyState uintptr
	if lpKeyState == nil {
		_lpKeyState = 0
	} else {
		_lpKeyState = uintptr(unsafe.Pointer(&lpKeyState[0]))
	}
	r1, _, err := procToAsciiEx.Call(uVirtKey, uScanCode, _lpKeyState, uintptr(unsafe.Pointer(&lpChar[0])), uFlags, dwhkl)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646320.aspx
func ToUnicode(wVirtKey uintptr, wScanCode uintptr, lpKeyState []uint8, pwszBuff []uint16, cchBuff uintptr, wFlags uintptr) (uintptr, error) {
	var _lpKeyState uintptr
	if lpKeyState == nil {
		_lpKeyState = 0
	} else {
		_lpKeyState = uintptr(unsafe.Pointer(&lpKeyState[0]))
	}
	r1, _, err := procToUnicode.Call(wVirtKey, wScanCode, _lpKeyState, uintptr(unsafe.Pointer(&pwszBuff[0])), cchBuff, wFlags)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646322.aspx
func ToUnicodeEx(wVirtKey uintptr, wScanCode uintptr, lpKeyState []uint8, pwszBuff []uint16, cchBuff uintptr, wFlags uintptr, dwhkl uintptr) (uintptr, error) {
	var _lpKeyState uintptr
	if lpKeyState == nil {
		_lpKeyState = 0
	} else {
		_lpKeyState = uintptr(unsafe.Pointer(&lpKeyState[0]))
	}
	r1, _, err := procToUnicodeEx.Call(wVirtKey, wScanCode, _lpKeyState, uintptr(unsafe.Pointer(&pwszBuff[0])), cchBuff, wFlags, dwhkl)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646324.aspx
func UnloadKeyboardLayout(hkl uintptr) (uintptr, error) {
	r1, _, err := procUnloadKeyboardLayout.Call(hkl)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646327.aspx
func UnregisterHotKey(hWnd uintptr, id uintptr) (uintptr, error) {
	r1, _, err := procUnregisterHotKey.Call(hWnd, id)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646329.aspx
func VkKeyScan(ch uintptr) (uintptr, error) {
	r1, _, err := procVkKeyScan.Call(ch)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646332.aspx
func VkKeyScanEx(ch uintptr, dwhkl uintptr) (uintptr, error) {
	r1, _, err := procVkKeyScanEx.Call(ch, dwhkl)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms646310.aspx
func sendInput(mode uintptr, nInputs uintptr, pInputs interface{}, cbSize uintptr) (uintptr, error) {
	var _pInputs uintptr
	switch mode {
	case _I_M:
		_pInputs = uintptr(unsafe.Pointer(&pInputs.([]INPUT_M)[0]))
	case _I_K:
		_pInputs = uintptr(unsafe.Pointer(&pInputs.([]INPUT_K)[0]))
	case _I_HW:
		_pInputs = uintptr(unsafe.Pointer(&pInputs.([]INPUT_HW)[0]))
	}
	r1, _, err := procSendInput.Call(nInputs, _pInputs, cbSize)
	return r1, err
}
