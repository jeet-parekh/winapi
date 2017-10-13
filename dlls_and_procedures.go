package winapi

import (
	"golang.org/x/sys/windows"
)

var (
	user32DLL = windows.NewLazySystemDLL("user32.dll")
)

var (
	// RAW INPUT
	// https://msdn.microsoft.com/en-us/library/ff468896.aspx
	procDefRawInputProc = user32DLL.NewProc("DefRawInputProc")
	procGetRawInputBuffer = user32DLL.NewProc("GetRawInputBuffer")
	procGetRawInputData = user32DLL.NewProc("GetRawInputData")
	procGetRawInputDeviceInfo = user32DLL.NewProc("GetRawInputDeviceInfoW")
	procGetRawInputDeviceList = user32DLL.NewProc("GetRawInputDeviceList")
	procGetRegisteredRawInputDevices = user32DLL.NewProc("GetRegisteredRawInputDevices")
	procRegisterRawInputDevices = user32DLL.NewProc("RegisterRawInputDevices")
)
