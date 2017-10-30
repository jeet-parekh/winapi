## Windows API wrappers for Go

---

### About `winapi`

- `winapi` provides wrappers around WinAPI functions.
- [Documentation: godoc](https://godoc.org/github.com/jeet-parekh/winapi)

APIs wrapped:

- [Configuration](https://msdn.microsoft.com/en-us/library/ff625300.aspx)
- [Hooks](https://msdn.microsoft.com/en-us/library/ms632589.aspx)
- [Keyboard Input](https://msdn.microsoft.com/en-us/library/ms645530.aspx)
- [Messages and Message Queues](https://msdn.microsoft.com/en-us/library/ms632590.aspx)
- [Mouse Input](https://msdn.microsoft.com/en-us/library/ms645533.aspx)
- [Multiple Document Interface](https://msdn.microsoft.com/en-us/library/ms632591.aspx)
- [Raw Input](https://msdn.microsoft.com/en-us/library/ms645536.aspx)
- [Timers](https://msdn.microsoft.com/en-us/library/ms632592.aspx)

---

WINAPI uses `union` which is not available in Go. To handle that, variants of structures and functions have been created for some APIs.

The structures and functions which have variants are mentioned below.

---

## RAW INPUT

`GetRawInputDeviceInfo` for uiCommand `RIDI_PREPARSEDDATA` isn't implemented.

### Abbreviations

|Abbreviation|Meaning|
|---|---|
|M|Mouse|
|MB|MouseButtons|
|MBI|MouseButtonInformation|
|K|Keyboard|
|HID|Human Interface Device|

### Structures

|WINAPI|Go|
|---|---|
|[RAWINPUT](https://msdn.microsoft.com/en-us/library/ms645562.aspx)|RAWINPUT\_MB<br>RAWINPUT\_MBI<br>RAWINPUT\_K<br>RAWINPUT\_HID|
|[RAWMOUSE](https://msdn.microsoft.com/en-us/library/ms645578.aspx)|RAWMOUSEBUTTONS<br>RAWMOUSEBUTTONINFO|
|[RID\_DEVICE\_INFO](https://msdn.microsoft.com/en-us/library/ms645581.aspx)|RIDI\_M<br>RIDI\_K<br>RIDI\_HID|

### Functions

|WINAPI|Go|
|---|---|
|[DefRawInputProc](https://msdn.microsoft.com/en-us/library/ms645594.aspx)|DefRawInputProcMB<br>DefRawInputProcMBI<br>DefRawInputProcK<br>DefRawInputProcHID|
|[GetRawInputBuffer](https://msdn.microsoft.com/en-us/library/ms645595.aspx)|GetRawInputBufferMB<br>GetRawInputBufferMBI<br>GetRawInputBufferK<br>GetRawInputBufferHID|
|[GetRawInputData](https://msdn.microsoft.com/en-us/library/ms645596.aspx)|GetRawInputDataMB<br>GetRawInputDataMBI<br>GetRawInputDataK<br>GetRawInputDataHID|
|[GetRawInputDeviceInfo](https://msdn.microsoft.com/en-us/library/ms645597.aspx)|GetRawInputDeviceName<br>GetRawInputDeviceInfoM<br>GetRawInputDeviceInfoK<br>GetRawInputDeviceInfoHID|

---

## KEYBOARD INPUT

### Abbreviations

|Abbreviation|Meaning|
|---|---|
|M|Mouse|
|K|Keyboard|
|HW|Hardware|

### Structures

|WINAPI|Go|
|---|---|
|[INPUT](https://msdn.microsoft.com/en-us/library/ms646270.aspx)|INPUT\_M<br>INPUT\_K<br>INPUT\_HW|

### Functions

|WINAPI|Go|
|---|---|
|[SendInput](https://msdn.microsoft.com/en-us/library/ms646310.aspx)|SendInputM<br>SendInputK<br>SendInputHW|

---
