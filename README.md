## Windows API wrappers for Go

---

APIs wrapped:

- [Raw Input](https://msdn.microsoft.com/en-us/library/ms645536.aspx)

---

WINAPI uses `union` which is not available in Go. To handle that, variants of structures and functions have been created for some APIs.

---

## RAW INPUT

### Abbreviations

|Abbreviation|Meaning|
|---|---|
|M|Mouse|
|MB|MouseButtons|
|MBI|MouseButtonInformation|
|K|Keyboard|
|HID|Human Interface Device|

MouseButtons (MB) and MouseButtonInformation (MBI) are variants of `RAWMOUSE` structure.

`GetRawInputDeviceInfo` for `uiCommand` `RIDI_PREPARSEDDATA` isn't implemented.

### Structures

|WINAPI|Go|
|---|---|
|[RAWHID](https://msdn.microsoft.com/en-us/library/ms645549.aspx)|RAWHID|
|[RAWINPUT](https://msdn.microsoft.com/en-us/library/ms645562.aspx)|RAWINPUT\_MB<br>RAWINPUT\_MBI<br>RAWINPUT\_K<br>RAWINPUT\_HID|
|[RAWINPUTDEVICE](https://msdn.microsoft.com/en-us/library/ms645565.aspx)|RAWINPUTDEVICE|
|[RAWINPUTDEVICELIST](https://msdn.microsoft.com/en-us/library/ms645568.aspx)|RAWINPUTDEVICELIST|
|[RAWINPUTHEADER](https://msdn.microsoft.com/en-us/library/ms645571.aspx)|RAWINPUTHEADER|
|[RAWKEYBOARD](https://msdn.microsoft.com/en-us/library/ms645575.aspx)|RAWKEYBOARD|
|[RAWMOUSE](https://msdn.microsoft.com/en-us/library/ms645578.aspx)|RAWMOUSEBUTTONS<br>RAWMOUSEBUTTONINFO|
|[RID\_DEVICE\_INFO](https://msdn.microsoft.com/en-us/library/ms645581.aspx)|RIDI\_M<br>RIDI\_K<br>RIDI\_HID|
|[RID\_DEVICE\_INFO\_HID](https://msdn.microsoft.com/en-us/library/ms645584.aspx)|RID\_DEVICE\_INFO\_HID|
|[RID\_DEVICE\_INFO\_KEYBOARD](https://msdn.microsoft.com/en-us/library/ms645587.aspx)|RID\_DEVICE\_INFO\_KEYBOARD|
|[RID\_DEVICE\_INFO\_MOUSE](https://msdn.microsoft.com/en-us/library/ms645589.aspx)|RID\_DEVICE\_INFO\_MOUSE|

### Functions

|WINAPI|Go|
|---|---|
|[DefRawInputProc](https://msdn.microsoft.com/en-us/library/ms645594.aspx)|DefRawInputProcMB<br>DefRawInputProcMBI<br>DefRawInputProcK<br>DefRawInputProcHID|
|[GetRawInputBuffer](https://msdn.microsoft.com/en-us/library/ms645595.aspx)|GetRawInputBufferMB<br>GetRawInputBufferMBI<br>GetRawInputBufferK<br>GetRawInputBufferHID|
|[GetRawInputData](https://msdn.microsoft.com/en-us/library/ms645596.aspx)|GetRawInputDataMB<br>GetRawInputDataMBI<br>GetRawInputDataK<br>GetRawInputDataHID|
|[GetRawInputDeviceInfo](https://msdn.microsoft.com/en-us/library/ms645597.aspx)|GetRawInputDeviceName<br>GetRawInputDeviceInfoM<br>GetRawInputDeviceInfoK<br>GetRawInputDeviceInfoHID|
|[GetRawInputDeviceList](https://msdn.microsoft.com/en-us/library/ms645598.aspx)|GetRawInputDeviceList|
|[GetRegisteredRawInputDevices](https://msdn.microsoft.com/en-us/library/ms645599.aspx)|GetRegisteredRawInputDevices|
|[RegisterRawInputDevices](https://msdn.microsoft.com/en-us/library/ms645600.aspx)|RegisterRawInputDevices|

---
