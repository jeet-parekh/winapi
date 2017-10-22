package winapi

import (
	"golang.org/x/sys/windows"
)

var (
	user32DLL = windows.NewLazySystemDLL("user32.dll")
	kernel32DLL = windows.NewLazySystemDLL("kernel32.dll")
)

var (
	// CONFIGURATION
	// https://msdn.microsoft.com/en-us/library/ff625301.aspx
	procGetSystemMetrics     = user32DLL.NewProc("GetSystemMetrics")
	procSystemParametersInfo = user32DLL.NewProc("SystemParametersInfoW")

	// HOOKS
	// https://msdn.microsoft.com/en-us/library/ff468842.aspx
	procCallMsgFilter       = user32DLL.NewProc("CallMsgFilterW")
	procCallNextHookEx      = user32DLL.NewProc("CallNextHookEx")
	procSetWindowsHookEx    = user32DLL.NewProc("SetWindowsHookExW")
	procUnhookWindowsHookEx = user32DLL.NewProc("UnhookWindowsHookEx")

	// KEYBOARD INPUT
	// https://msdn.microsoft.com/en-us/library/ff468859.aspx
	procActivateKeyboardLayout = user32DLL.NewProc("ActivateKeyboardLayout")
	procBlockInput             = user32DLL.NewProc("BlockInput")
	procEnableWindow           = user32DLL.NewProc("EnableWindow")
	procGetActiveWindow        = user32DLL.NewProc("GetActiveWindow")
	procGetAsyncKeyState       = user32DLL.NewProc("GetAsyncKeyState")
	procGetFocus               = user32DLL.NewProc("GetFocus")
	procGetKBCodePage          = user32DLL.NewProc("GetKBCodePage")
	procGetKeyboardLayout      = user32DLL.NewProc("GetKeyboardLayout")
	procGetKeyboardLayoutList  = user32DLL.NewProc("GetKeyboardLayoutList")
	procGetKeyboardLayoutName  = user32DLL.NewProc("GetKeyboardLayoutNameW")
	procGetKeyboardState       = user32DLL.NewProc("GetKeyboardState")
	procGetKeyboardType        = user32DLL.NewProc("GetKeyboardType")
	procGetKeyNameText         = user32DLL.NewProc("GetKeyNameTextW")
	procGetKeyState            = user32DLL.NewProc("GetKeyState")
	procGetLastInputInfo       = user32DLL.NewProc("GetLastInputInfo")
	procIsWindowEnabled        = user32DLL.NewProc("IsWindowEnabled")
	procKeybd_event            = user32DLL.NewProc("keybd_event")
	procLoadKeyboardLayout     = user32DLL.NewProc("LoadKeyboardLayoutW")
	procMapVirtualKey          = user32DLL.NewProc("MapVirtualKeyW")
	procMapVirtualKeyEx        = user32DLL.NewProc("MapVirtualKeyExW")
	procOemKeyScan             = user32DLL.NewProc("OemKeyScan")
	procRegisterHotKey         = user32DLL.NewProc("RegisterHotKey")
	procSendInput              = user32DLL.NewProc("SendInput")
	procSetActiveWindow        = user32DLL.NewProc("SetActiveWindow")
	procSetFocus               = user32DLL.NewProc("SetFocus")
	procSetKeyboardState       = user32DLL.NewProc("SetKeyboardState")
	procToAscii                = user32DLL.NewProc("ToAscii")
	procToAsciiEx              = user32DLL.NewProc("ToAsciiEx")
	procToUnicode              = user32DLL.NewProc("ToUnicode")
	procToUnicodeEx            = user32DLL.NewProc("ToUnicodeEx")
	procUnloadKeyboardLayout   = user32DLL.NewProc("UnloadKeyboardLayout")
	procUnregisterHotKey       = user32DLL.NewProc("UnregisterHotKey")
	procVkKeyScan              = user32DLL.NewProc("VkKeyScanW")
	procVkKeyScanEx            = user32DLL.NewProc("VkKeyScanExW")

	// MESSAGES AND MESSAGE QUEUES
	// https://msdn.microsoft.com/en-us/library/ff468870.aspx
	procBroadcastSystemMessage   = user32DLL.NewProc("BroadcastSystemMessageW")
	procBroadcastSystemMessageEx = user32DLL.NewProc("BroadcastSystemMessageExW")
	procDispatchMessage          = user32DLL.NewProc("DispatchMessageW")
	procGetInputState            = user32DLL.NewProc("GetInputState")
	procGetMessage               = user32DLL.NewProc("GetMessageW")
	procGetMessageExtraInfo      = user32DLL.NewProc("GetMessageExtraInfo")
	procGetMessagePos            = user32DLL.NewProc("GetMessagePos")
	procGetMessageTime           = user32DLL.NewProc("GetMessageTime")
	procGetQueueStatus           = user32DLL.NewProc("GetQueueStatus")
	procInSendMessage            = user32DLL.NewProc("InSendMessage")
	procInSendMessageEx          = user32DLL.NewProc("InSendMessageEx")
	procPeekMessage              = user32DLL.NewProc("PeekMessageW")
	procPostMessage              = user32DLL.NewProc("PostMessageW")
	procPostQuitMessage          = user32DLL.NewProc("PostQuitMessage")
	procPostThreadMessage        = user32DLL.NewProc("PostThreadMessageW")
	procRegisterWindowMessage    = user32DLL.NewProc("RegisterWindowMessageW")
	procReplyMessage             = user32DLL.NewProc("ReplyMessage")
	procSendMessage              = user32DLL.NewProc("SendMessageW")
	procSendMessageCallback      = user32DLL.NewProc("SendMessageCallbackW")
	procSendMessageTimeout       = user32DLL.NewProc("SendMessageTimeoutW")
	procSendNotifyMessage        = user32DLL.NewProc("SendNotifyMessageW")
	procSetMessageExtraInfo      = user32DLL.NewProc("SetMessageExtraInfo")
	procTranslateMessage         = user32DLL.NewProc("TranslateMessage")
	procWaitMessage              = user32DLL.NewProc("WaitMessage")

	// MULTIPLE DOCUMENT INTERFACE
	// https://msdn.microsoft.com/en-us/library/ff468891.aspx
	procCreateMDIWindow      = user32DLL.NewProc("CreateMDIWindowW")
	procDefFrameProc         = user32DLL.NewProc("DefFrameProcW")
	procDefMDIChildProc      = user32DLL.NewProc("DefMDIChildProcW")
	procTranslateMDISysAccel = user32DLL.NewProc("TranslateMDISysAccel")

	// MOUSE INPUT
	// https://msdn.microsoft.com/en-us/library/ff468875.aspx
	procDragDetect           = user32DLL.NewProc("DragDetect")
	procEnableMouseInPointer = user32DLL.NewProc("EnableMouseInPointer")
	procGetCapture           = user32DLL.NewProc("GetCapture")
	procGetDoubleClickTime   = user32DLL.NewProc("GetDoubleClickTime")
	procGetMouseMovePointsEx = user32DLL.NewProc("GetMouseMovePointsEx")
	procMouse_event          = user32DLL.NewProc("mouse_event")
	procReleaseCapture       = user32DLL.NewProc("ReleaseCapture")
	procSetCapture           = user32DLL.NewProc("SetCapture")
	procSetDoubleClickTime   = user32DLL.NewProc("SetDoubleClickTime")
	procSwapMouseButton      = user32DLL.NewProc("SwapMouseButton")
	procTrackMouseEvent      = user32DLL.NewProc("TrackMouseEvent")

	// RAW INPUT
	// https://msdn.microsoft.com/en-us/library/ff468896.aspx
	procDefRawInputProc              = user32DLL.NewProc("DefRawInputProc")
	procGetRawInputBuffer            = user32DLL.NewProc("GetRawInputBuffer")
	procGetRawInputData              = user32DLL.NewProc("GetRawInputData")
	procGetRawInputDeviceInfo        = user32DLL.NewProc("GetRawInputDeviceInfoW")
	procGetRawInputDeviceList        = user32DLL.NewProc("GetRawInputDeviceList")
	procGetRegisteredRawInputDevices = user32DLL.NewProc("GetRegisteredRawInputDevices")
	procRegisterRawInputDevices      = user32DLL.NewProc("RegisterRawInputDevices")

	// TIMERS
	// https://msdn.microsoft.com/en-us/library/ff468913.aspx
	procKillTimer           = user32DLL.NewProc("KillTimer")
	procSetCoalescableTimer = user32DLL.NewProc("SetCoalescableTimer")
	procSetTimer            = user32DLL.NewProc("SetTimer")

	// PROCESSES AND THREADS
	// https://msdn.microsoft.com/en-us/library/ms684847.aspx
	procGetCurrentThreadId = kernel32DLL.NewProc("GetCurrentThreadId")

	// CONSOLE
	// https://docs.microsoft.com/en-us/windows/console/console-functions
	procFlushConsoleInputBuffer = kernel32DLL.NewProc("FlushConsoleInputBuffer")
	procGetConsoleMode = kernel32DLL.NewProc("GetConsoleMode")
	procGetStdHandle = kernel32DLL.NewProc("GetStdHandle")
	procPeekConsoleInput = kernel32DLL.NewProc("PeekConsoleInputW")
	procReadConsoleInput = kernel32DLL.NewProc("ReadConsoleInputW")
	procSetConsoleMode = kernel32DLL.NewProc("SetConsoleMode")
)
