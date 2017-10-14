package winapi

// https://msdn.microsoft.com/en-us/library/ms644903.aspx
func KillTimer(hWnd uintptr, uIDEvent uintptr) (uintptr, error) {
	r1, _, err := procKillTimer.Call(hWnd, uIDEvent)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/hh405404.aspx
func SetCoalescableTimer(hwnd uintptr, nIDEvent uintptr, uElapse uintptr, lpTimerFunc uintptr, uToleranceDelay uintptr) (uintptr, error) {
	r1, _, err := procSetCoalescableTimer.Call(hwnd, nIDEvent, uElapse, lpTimerFunc, uToleranceDelay)
	return r1, err
}

// https://msdn.microsoft.com/en-us/library/ms644906.aspx
func SetTimer(hWnd uintptr, nIDEvent uintptr, uElapse uintptr, lpTimerFunc uintptr) (uintptr, error) {
	r1, _, err := procSetTimer.Call(hWnd, nIDEvent, uElapse, lpTimerFunc)
	return r1, err
}
