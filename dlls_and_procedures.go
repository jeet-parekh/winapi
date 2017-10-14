package winapi

import (
	"golang.org/x/sys/windows"
)

var (
	user32DLL = windows.NewLazySystemDLL("user32.dll")
)

var (
	// CONFIGURATION
	// https://msdn.microsoft.com/en-us/library/ff625301.aspx
	procGetSystemMetrics = user32DLL.NewProc("GetSystemMetrics")
	procSystemParametersInfo = user32DLL.NewProc("SystemParametersInfoW")
	
	// HOOKS
	// https://msdn.microsoft.com/en-us/library/ff468842.aspx
	procCallMsgFilter = user32DLL.NewProc("CallMsgFilterW")
	procCallNextHookEx = user32DLL.NewProc("CallNextHookEx")
	procSetWindowsHookEx = user32DLL.NewProc("SetWindowsHookExW")
	procUnhookWindowsHookEx = user32DLL.NewProc("UnhookWindowsHookEx")

	// MESSAGES AND MESSAGE QUEUES
	// https://msdn.microsoft.com/en-us/library/ff468870.aspx
	procBroadcastSystemMessage = user32DLL.NewProc("BroadcastSystemMessageW")
	procBroadcastSystemMessageEx = user32DLL.NewProc("BroadcastSystemMessageExW")
	procDispatchMessage = user32DLL.NewProc("DispatchMessageW")
	procGetInputState = user32DLL.NewProc("GetInputState")
	procGetMessage = user32DLL.NewProc("GetMessageW")
	procGetMessageExtraInfo = user32DLL.NewProc("GetMessageExtraInfo")
	procGetMessagePos = user32DLL.NewProc("GetMessagePos")
	procGetMessageTime = user32DLL.NewProc("GetMessageTime")
	procGetQueueStatus = user32DLL.NewProc("GetQueueStatus")
	procInSendMessage = user32DLL.NewProc("InSendMessage")
	procInSendMessageEx = user32DLL.NewProc("InSendMessageEx")
	procPeekMessage = user32DLL.NewProc("PeekMessageW")
	procPostMessage = user32DLL.NewProc("PostMessageW")
	procPostQuitMessage = user32DLL.NewProc("PostQuitMessage")
	procPostThreadMessage = user32DLL.NewProc("PostThreadMessageW")
	procRegisterWindowMessage = user32DLL.NewProc("RegisterWindowMessageW")
	procReplyMessage = user32DLL.NewProc("ReplyMessage")
	procSendMessage = user32DLL.NewProc("SendMessageW")
	procSendMessageCallback = user32DLL.NewProc("SendMessageCallbackW")
	procSendMessageTimeout = user32DLL.NewProc("SendMessageTimeoutW")
	procSendNotifyMessage = user32DLL.NewProc("SendNotifyMessageW")
	procSetMessageExtraInfo = user32DLL.NewProc("SetMessageExtraInfo")
	procTranslateMessage = user32DLL.NewProc("TranslateMessage")
	procWaitMessage = user32DLL.NewProc("WaitMessage")

	// MULTIPLE DOCUMENT INTERFACE
	// https://msdn.microsoft.com/en-us/library/ff468891.aspx
	procCreateMDIWindow = user32DLL.NewProc("CreateMDIWindowW")
	procDefFrameProc = user32DLL.NewProc("DefFrameProcW")
	procDefMDIChildProc = user32DLL.NewProc("DefMDIChildProcW")
	procTranslateMDISysAccel = user32DLL.NewProc("TranslateMDISysAccel")

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
