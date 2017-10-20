package winapi

// https://msdn.microsoft.com/en-us/library/ms683183.aspx
func GetCurrentThreadId() (uintptr, error) {
	r1, _, err := procGetCurrentThreadId.Call()
	return r1, err
}
