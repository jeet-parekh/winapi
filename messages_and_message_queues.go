package winapi

import (
	"unsafe"
)

// https://msdn.microsoft.com/en-us/library/ms644932.aspx
func BroadcastSystemMessage(dwFlags uintptr, lpdwRecipients *uintptr, uiMessage uintptr, wParam uintptr, lParam uintptr) (uintptr, error) {
	r1, _, err := procBroadcastSystemMessage.Call(dwFlags, uintptr(unsafe.Pointer(lpdwRecipients)), uiMessage, wParam, lParam)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644933.aspx
func BroadcastSystemMessageEx(dwFlags uintptr, lpdwRecipients *uintptr, uiMessage uintptr, wParam uintptr, lParam uintptr, pBSMInfo *BSMINFO) (uintptr, error) {
	r1, _, err := procBroadcastSystemMessageEx.Call(dwFlags, uintptr(unsafe.Pointer(lpdwRecipients)), uiMessage, wParam, lParam, uintptr(unsafe.Pointer(pBSMInfo)))
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644934.aspx
func DispatchMessage(lpMsg *MSG) (uintptr, error) {
	r1, _, err := procDispatchMessage.Call(uintptr(unsafe.Pointer(lpMsg)))
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644935.aspx
func GetInputState() (uintptr, error) {
	r1, _, err := procGetInputState.Call()
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644936.aspx
func GetMessage(lpMsg *MSG, hWnd uintptr, wMsgFilterMin uintptr, wMsgFilterMax uintptr) (uintptr, error) {
	r1, _, err := procGetMessage.Call(uintptr(unsafe.Pointer(lpMsg)), hWnd, wMsgFilterMin, wMsgFilterMax)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644937.aspx
func GetMessageExtraInfo() (uintptr, error) {
	r1, _, err := procGetMessageExtraInfo.Call()
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644938.aspx
func GetMessagePos() (uintptr, error) {
	r1, _, err := procGetMessagePos.Call()
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644939.aspx
func GetGetMessageTime() (uintptr, error) {
	r1, _, err := procGetMessageTime.Call()
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644940.aspx
func GetQueueStatus(flags uintptr) (uintptr, error) {
	r1, _, err := procGetQueueStatus.Call(flags)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644941.aspx
func InSendMessage() (uintptr, error) {
	r1, _, err := procInSendMessage.Call()
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644942.aspx
func InSendMessageEx() (uintptr, error) {
	r1, _, err := procInSendMessageEx.Call()
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644943.aspx
func PeekMessage(lpMsg *MSG, hWnd uintptr, wMsgFilterMin uintptr, wMsgFilterMax uintptr, wRemoveMsg uintptr) (uintptr, error) {
	r1, _, err := procPeekMessage.Call(uintptr(unsafe.Pointer(lpMsg)), hWnd, wMsgFilterMin, wMsgFilterMax, wRemoveMsg)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644944.aspx
func PostMessage(hWnd uintptr, Msg uintptr, wParam uintptr, lParam uintptr) (uintptr, error) {
	r1, _, err := procPostMessage.Call(hWnd, Msg, wParam, lParam)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644945.aspx
func PostQuitMessage(nExitCode int32) (uintptr, error) {
	r1, _, err := procPostQuitMessage.Call(int32ToUintptr(nExitCode))
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644946.aspx
func PostThreadMessage(idThread uintptr, Msg uintptr, wParam uintptr, lParam uintptr) (uintptr, error) {
	r1, _, err := procPostThreadMessage.Call(idThread, Msg, wParam, lParam)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644947.aspx
func RegisterWindowMessage(lpString string) (uintptr, error) {
	_lpString, err := strToUintptr(lpString)
	if err != nil {
		return 0, err
	}
	r1, _, err := procRegisterWindowMessage.Call(uintptr(unsafe.Pointer(_lpString)))
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644948.aspx
func ReplyMessage(lResult uintptr) (uintptr, error) {
	r1, _, err := procReplyMessage.Call(lResult)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644950.aspx
func SendMessage(hWnd uintptr, Msg uintptr, wParam uintptr, lParam uintptr) (uintptr, error) {
	r1, _, err := procSendMessage.Call(hWnd, Msg, wParam, lParam)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644951.aspx
func SendMessageCallback(hWnd uintptr, Msg uintptr, wParam uintptr, lParam uintptr, lpCallback uintptr, dwData uintptr) (uintptr, error) {
	r1, _, err := procSendMessageCallback.Call(hWnd, Msg, wParam, lParam, lpCallback, dwData)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644952.aspx
func SendMessageTimeout(hWnd uintptr, Msg uintptr, wParam uintptr, lParam uintptr, fuFlags uintptr, uTimeout uintptr, lpdwResult *uintptr) (uintptr, error) {
	r1, _, err := procSendMessageTimeout.Call(hWnd, Msg, wParam, lParam, fuFlags, uTimeout, uintptr(unsafe.Pointer(lpdwResult)))
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644953.aspx
func SendNotifyMessage(hWnd uintptr, Msg uintptr, wParam uintptr, lParam uintptr) (uintptr, error) {
	r1, _, err := procSendNotifyMessage.Call(hWnd, Msg, wParam, lParam)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644954.aspx
func SetMessageExtraInfo(lParam uintptr) (uintptr, error) {
	r1, _, err := procSetMessageExtraInfo.Call(lParam)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644955.aspx
func TranslateMessage(lpMsg *MSG) (uintptr, error) {
	r1, _, err := procTranslateMessage.Call(uintptr(unsafe.Pointer(lpMsg)))
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644956.aspx
func WaitMessage() (uintptr, error) {
	r1, _, err := procWaitMessage.Call()
	return r1, err
}
