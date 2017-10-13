package winapi

import (
	"unsafe"
)

// https://msdn.microsoft.com/en-us/library/ms645594.aspx
func DefRawInputProcMB(paRawInput []RAWINPUT_MB, nInput uintptr, cbSizeHeader uintptr) (uintptr, error) {
	return defRawInputProc(_RI_MB, paRawInput, nInput, cbSizeHeader)
}

// https://msdn.microsoft.com/en-us/library/ms645594.aspx
func DefRawInputProcMBI(paRawInput []RAWINPUT_MBI, nInput uintptr, cbSizeHeader uintptr) (uintptr, error) {
	return defRawInputProc(_RI_MBI, paRawInput, nInput, cbSizeHeader)
}

// https://msdn.microsoft.com/en-us/library/ms645594.aspx
func DefRawInputProcK(paRawInput []RAWINPUT_K, nInput uintptr, cbSizeHeader uintptr) (uintptr, error) {
	return defRawInputProc(_RI_K, paRawInput, nInput, cbSizeHeader)
}

// https://msdn.microsoft.com/en-us/library/ms645594.aspx
func DefRawInputProcHID(paRawInput []RAWINPUT_HID, nInput uintptr, cbSizeHeader uintptr) (uintptr, error) {
	return defRawInputProc(_RI_HID, paRawInput, nInput, cbSizeHeader)
}

// https://msdn.microsoft.com/en-us/library/ms645595.aspx
func GetRawInputBufferMB(pData []RAWINPUT_MB, pcbSize *uintptr, cbSizeHeader uintptr) (uintptr, error) {
	return getRawInputBuffer(_RI_MB, pData, pcbSize, cbSizeHeader)
}

// https://msdn.microsoft.com/en-us/library/ms645595.aspx
func GetRawInputBufferMBI(pData []RAWINPUT_MBI, pcbSize *uintptr, cbSizeHeader uintptr) (uintptr, error) {
	return getRawInputBuffer(_RI_MBI, pData, pcbSize, cbSizeHeader)
}

// https://msdn.microsoft.com/en-us/library/ms645595.aspx
func GetRawInputBufferK(pData []RAWINPUT_K, pcbSize *uintptr, cbSizeHeader uintptr) (uintptr, error) {
	return getRawInputBuffer(_RI_K, pData, pcbSize, cbSizeHeader)
}

// https://msdn.microsoft.com/en-us/library/ms645595.aspx
func GetRawInputBufferHID(pData []RAWINPUT_HID, pcbSize *uintptr, cbSizeHeader uintptr) (uintptr, error) {
	return getRawInputBuffer(_RI_HID, pData, pcbSize, cbSizeHeader)
}

// https://msdn.microsoft.com/en-us/library/ms645596.aspx
func GetRawInputDataMB(hRawInput uintptr, uiCommand uintptr, pData *RAWINPUT_MB, pcbSize *uintptr, cbSizeHeader uintptr) (uintptr, error) {
	return getRawInputData(_RI_MB, hRawInput, uiCommand, pData, pcbSize, cbSizeHeader)
}

// https://msdn.microsoft.com/en-us/library/ms645596.aspx
func GetRawInputDataMBI(hRawInput uintptr, uiCommand uintptr, pData *RAWINPUT_MBI, pcbSize *uintptr, cbSizeHeader uintptr) (uintptr, error) {
	return getRawInputData(_RI_MBI, hRawInput, uiCommand, pData, pcbSize, cbSizeHeader)
}

// https://msdn.microsoft.com/en-us/library/ms645596.aspx
func GetRawInputDataK(hRawInput uintptr, uiCommand uintptr, pData *RAWINPUT_K, pcbSize *uintptr, cbSizeHeader uintptr) (uintptr, error) {
	return getRawInputData(_RI_K, hRawInput, uiCommand, pData, pcbSize, cbSizeHeader)
}

// https://msdn.microsoft.com/en-us/library/ms645596.aspx
func GetRawInputDataHID(hRawInput uintptr, uiCommand uintptr, pData *RAWINPUT_HID, pcbSize *uintptr, cbSizeHeader uintptr) (uintptr, error) {
	return getRawInputData(_RI_HID, hRawInput, uiCommand, pData, pcbSize, cbSizeHeader)
}

// https://msdn.microsoft.com/en-us/library/ms645597.aspx
func GetRawInputDeviceName(hDevice uintptr, pData []uint16, pcbSize *uintptr) (uintptr, error) {
	return getRawInputDeviceInfo(0, hDevice, RIDI_DEVICENAME, pData, pcbSize)
}

// https://msdn.microsoft.com/en-us/library/ms645597.aspx
func GetRawInputDeviceInfoM(hDevice uintptr, pData *RIDI_M, pcbSize *uintptr) (uintptr, error) {
	return getRawInputDeviceInfo(RIM_TYPEMOUSE, hDevice, RIDI_DEVICEINFO, pData, pcbSize)
}

// https://msdn.microsoft.com/en-us/library/ms645597.aspx
func GetRawInputDeviceInfoK(hDevice uintptr, pData *RIDI_K, pcbSize *uintptr) (uintptr, error) {
	return getRawInputDeviceInfo(RIM_TYPEKEYBOARD, hDevice, RIDI_DEVICEINFO, pData, pcbSize)
}

// https://msdn.microsoft.com/en-us/library/ms645597.aspx
func GetRawInputDeviceInfoHID(hDevice uintptr, pData *RIDI_HID, pcbSize *uintptr) (uintptr, error) {
	return getRawInputDeviceInfo(RIM_TYPEHID, hDevice, RIDI_DEVICEINFO, pData, pcbSize)
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

// https://msdn.microsoft.com/en-us/library/ms645594.aspx
func defRawInputProc(mode uintptr, paRawInput interface{}, nInput uintptr, cbSizeHeader uintptr) (uintptr, error) {
	var _paRawInput uintptr
	switch mode {
	case _RI_MB:
		_paRawInput = uintptr(unsafe.Pointer(&paRawInput.([]RAWINPUT_MB)[0]))
	case _RI_MBI:
		_paRawInput = uintptr(unsafe.Pointer(&paRawInput.([]RAWINPUT_MBI)[0]))
	case _RI_K:
		_paRawInput = uintptr(unsafe.Pointer(&paRawInput.([]RAWINPUT_K)[0]))
	case _RI_HID:
		_paRawInput = uintptr(unsafe.Pointer(&paRawInput.([]RAWINPUT_HID)[0]))
	}
	r1, _, err := procDefRawInputProc.Call(_paRawInput, nInput, cbSizeHeader)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms645595.aspx
func getRawInputBuffer(mode uintptr, pData interface{}, pcbSize *uintptr, cbSizeHeader uintptr) (uintptr, error) {
	var _pData uintptr
	if pData == nil {
		_pData = 0
	} else {
		switch mode {
		case _RI_MB:
			_pData = uintptr(unsafe.Pointer(&pData.([]RAWINPUT_MB)[0]))
		case _RI_MBI:
			_pData = uintptr(unsafe.Pointer(&pData.([]RAWINPUT_MBI)[0]))
		case _RI_K:
			_pData = uintptr(unsafe.Pointer(&pData.([]RAWINPUT_K)[0]))
		case _RI_HID:
			_pData = uintptr(unsafe.Pointer(&pData.([]RAWINPUT_HID)[0]))
		}
	}
	r1, _, err := procGetRawInputBuffer.Call(_pData, uintptr(unsafe.Pointer(pcbSize)), cbSizeHeader)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms645596.aspx
func getRawInputData(mode uintptr, hRawInput uintptr, uiCommand uintptr, pData interface{}, pcbSize *uintptr, cbSizeHeader uintptr) (uintptr, error) {
	var _pData uintptr
	if pData == nil {
		_pData = 0
	} else {
		switch mode {
		case _RI_MB:
			_pData = uintptr(unsafe.Pointer(pData.(*RAWINPUT_MB)))
		case _RI_MBI:
			_pData = uintptr(unsafe.Pointer(pData.(*RAWINPUT_MBI)))
		case _RI_K:
			_pData = uintptr(unsafe.Pointer(pData.(*RAWINPUT_K)))
		case _RI_HID:
			_pData = uintptr(unsafe.Pointer(pData.(*RAWINPUT_HID)))
		}
	}
	r1, _, err := procGetRawInputData.Call(hRawInput, uiCommand, _pData, uintptr(unsafe.Pointer(pcbSize)), cbSizeHeader)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms645597.aspx
func getRawInputDeviceInfo(mode uintptr, hDevice uintptr, uiCommand uintptr, pData interface{}, pcbSize *uintptr) (uintptr, error) {
	var _pData uintptr
	if pData == nil {
		_pData = 0
	} else if uiCommand == RIDI_DEVICENAME {
		_pData = uintptr(unsafe.Pointer(&pData.([]uint16)[0]))
	} else if uiCommand == RIDI_DEVICEINFO {
		switch mode {
		case RIM_TYPEMOUSE:
			_pData = uintptr(unsafe.Pointer(pData.(*RIDI_M)))
		case RIM_TYPEKEYBOARD:
			_pData = uintptr(unsafe.Pointer(pData.(*RIDI_K)))
		case RIM_TYPEHID:
			_pData = uintptr(unsafe.Pointer(pData.(*RIDI_HID)))
		}
	}
	r1, _, err := procGetRawInputDeviceInfo.Call(hDevice, uiCommand, _pData, uintptr(unsafe.Pointer(pcbSize)))
	return r1, err
}
