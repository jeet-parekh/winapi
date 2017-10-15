package winapi

import (
	"unsafe"
)

// https://msdn.microsoft.com/en-us/library/ms645594.aspx
func DefRawInputProcMB(paRawInput []RAWINPUT_MB, nInput uintptr, cbSizeHeader uintptr) (uintptr, error) {
	r1, _, err := procDefRawInputProc.Call(uintptr(unsafe.Pointer(&paRawInput[0])), nInput, cbSizeHeader)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms645594.aspx
func DefRawInputProcMBI(paRawInput []RAWINPUT_MBI, nInput uintptr, cbSizeHeader uintptr) (uintptr, error) {
	r1, _, err := procDefRawInputProc.Call(uintptr(unsafe.Pointer(&paRawInput[0])), nInput, cbSizeHeader)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms645594.aspx
func DefRawInputProcK(paRawInput []RAWINPUT_K, nInput uintptr, cbSizeHeader uintptr) (uintptr, error) {
	r1, _, err := procDefRawInputProc.Call(uintptr(unsafe.Pointer(&paRawInput[0])), nInput, cbSizeHeader)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms645594.aspx
func DefRawInputProcHID(paRawInput []RAWINPUT_HID, nInput uintptr, cbSizeHeader uintptr) (uintptr, error) {
	r1, _, err := procDefRawInputProc.Call(uintptr(unsafe.Pointer(&paRawInput[0])), nInput, cbSizeHeader)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms645595.aspx
func GetRawInputBufferMB(pData []RAWINPUT_MB, pcbSize *uintptr, cbSizeHeader uintptr) (uintptr, error) {
	var _pData uintptr
	if pData == nil {
		_pData = 0
	} else {
		_pData = uintptr(unsafe.Pointer(&pData[0]))
	}
	r1, _, err := procGetRawInputBuffer.Call(_pData, uintptr(unsafe.Pointer(pcbSize)), cbSizeHeader)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms645595.aspx
func GetRawInputBufferMBI(pData []RAWINPUT_MBI, pcbSize *uintptr, cbSizeHeader uintptr) (uintptr, error) {
	var _pData uintptr
	if pData == nil {
		_pData = 0
	} else {
		_pData = uintptr(unsafe.Pointer(&pData[0]))
	}
	r1, _, err := procGetRawInputBuffer.Call(_pData, uintptr(unsafe.Pointer(pcbSize)), cbSizeHeader)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms645595.aspx
func GetRawInputBufferK(pData []RAWINPUT_K, pcbSize *uintptr, cbSizeHeader uintptr) (uintptr, error) {
	var _pData uintptr
	if pData == nil {
		_pData = 0
	} else {
		_pData = uintptr(unsafe.Pointer(&pData[0]))
	}
	r1, _, err := procGetRawInputBuffer.Call(_pData, uintptr(unsafe.Pointer(pcbSize)), cbSizeHeader)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms645595.aspx
func GetRawInputBufferHID(pData []RAWINPUT_HID, pcbSize *uintptr, cbSizeHeader uintptr) (uintptr, error) {
	var _pData uintptr
	if pData == nil {
		_pData = 0
	} else {
		_pData = uintptr(unsafe.Pointer(&pData[0]))
	}
	r1, _, err := procGetRawInputBuffer.Call(_pData, uintptr(unsafe.Pointer(pcbSize)), cbSizeHeader)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms645596.aspx
func GetRawInputDataMB(hRawInput uintptr, uiCommand uintptr, pData *RAWINPUT_MB, pcbSize *uintptr, cbSizeHeader uintptr) (uintptr, error) {
	var _pData uintptr
	if pData == nil {
		_pData = 0
	} else {
		_pData = uintptr(unsafe.Pointer(pData))
	}
	r1, _, err := procGetRawInputData.Call(hRawInput, uiCommand, _pData, uintptr(unsafe.Pointer(pcbSize)), cbSizeHeader)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms645596.aspx
func GetRawInputDataMBI(hRawInput uintptr, uiCommand uintptr, pData *RAWINPUT_MBI, pcbSize *uintptr, cbSizeHeader uintptr) (uintptr, error) {
	var _pData uintptr
	if pData == nil {
		_pData = 0
	} else {
		_pData = uintptr(unsafe.Pointer(pData))
	}
	r1, _, err := procGetRawInputData.Call(hRawInput, uiCommand, _pData, uintptr(unsafe.Pointer(pcbSize)), cbSizeHeader)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms645596.aspx
func GetRawInputDataK(hRawInput uintptr, uiCommand uintptr, pData *RAWINPUT_K, pcbSize *uintptr, cbSizeHeader uintptr) (uintptr, error) {
	var _pData uintptr
	if pData == nil {
		_pData = 0
	} else {
		_pData = uintptr(unsafe.Pointer(pData))
	}
	r1, _, err := procGetRawInputData.Call(hRawInput, uiCommand, _pData, uintptr(unsafe.Pointer(pcbSize)), cbSizeHeader)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms645596.aspx
func GetRawInputDataHID(hRawInput uintptr, uiCommand uintptr, pData *RAWINPUT_HID, pcbSize *uintptr, cbSizeHeader uintptr) (uintptr, error) {
	var _pData uintptr
	if pData == nil {
		_pData = 0
	} else {
		_pData = uintptr(unsafe.Pointer(pData))
	}
	r1, _, err := procGetRawInputData.Call(hRawInput, uiCommand, _pData, uintptr(unsafe.Pointer(pcbSize)), cbSizeHeader)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms645597.aspx
func GetRawInputDeviceName(hDevice uintptr, pData []uint16, pcbSize *uintptr) (uintptr, error) {
	var _pData uintptr
	if pData == nil {
		_pData = 0
	} else {
		_pData = uintptr(unsafe.Pointer(&pData[0]))
	}
	r1, _, err := procGetRawInputDeviceInfo.Call(hDevice, RIDI_DEVICENAME, _pData, uintptr(unsafe.Pointer(pcbSize)))
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms645597.aspx
func GetRawInputDeviceInfoM(hDevice uintptr, pData *RIDI_M, pcbSize *uintptr) (uintptr, error) {
	var _pData uintptr
	if pData == nil {
		_pData = 0
	} else {
		_pData = uintptr(unsafe.Pointer(pData))
	}
	r1, _, err := procGetRawInputDeviceInfo.Call(hDevice, RIDI_DEVICEINFO, _pData, uintptr(unsafe.Pointer(pcbSize)))
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms645597.aspx
func GetRawInputDeviceInfoK(hDevice uintptr, pData *RIDI_K, pcbSize *uintptr) (uintptr, error) {
	var _pData uintptr
	if pData == nil {
		_pData = 0
	} else {
		_pData = uintptr(unsafe.Pointer(pData))
	}
	r1, _, err := procGetRawInputDeviceInfo.Call(hDevice, RIDI_DEVICEINFO, _pData, uintptr(unsafe.Pointer(pcbSize)))
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms645597.aspx
func GetRawInputDeviceInfoHID(hDevice uintptr, pData *RIDI_HID, pcbSize *uintptr) (uintptr, error) {
	var _pData uintptr
	if pData == nil {
		_pData = 0
	} else {
		_pData = uintptr(unsafe.Pointer(pData))
	}
	r1, _, err := procGetRawInputDeviceInfo.Call(hDevice, RIDI_DEVICEINFO, _pData, uintptr(unsafe.Pointer(pcbSize)))
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms645598.aspx
func GetRawInputDeviceList(pRawInputDeviceList []RAWINPUTDEVICELIST, puiNumDevices *uintptr, cbSize uintptr) (uintptr, error) {
	var _pRawInputDeviceList uintptr
	if pRawInputDeviceList == nil {
		_pRawInputDeviceList = 0
	} else {
		_pRawInputDeviceList = uintptr(unsafe.Pointer(&pRawInputDeviceList[0]))
	}
	r1, _, err := procGetRawInputDeviceList.Call(_pRawInputDeviceList, uintptr(unsafe.Pointer(puiNumDevices)), cbSize)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms645599.aspx
func GetRegisteredRawInputDevices(pRawInputDevices []RAWINPUTDEVICE, puiNumDevices *uintptr, cbSize uintptr) (uintptr, error) {
	var _pRawInputDevices uintptr
	if pRawInputDevices == nil {
		_pRawInputDevices = 0
	} else {
		_pRawInputDevices = uintptr(unsafe.Pointer(&pRawInputDevices[0]))
	}
	r1, _, err := procGetRegisteredRawInputDevices.Call(_pRawInputDevices, uintptr(unsafe.Pointer(puiNumDevices)), cbSize)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms645600.aspx
func RegisterRawInputDevices(pRawInputDevices []RAWINPUTDEVICE, uiNumDevices uintptr, cbSize uintptr) (uintptr, error) {
	r1, _, err := procRegisterRawInputDevices.Call(uintptr(unsafe.Pointer(&pRawInputDevices[0])), uiNumDevices, cbSize)
	return r1, err
}
