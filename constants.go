package winapi

const (
	RIM_INPUT     uintptr = 0
	RIM_INPUTSINK uintptr = 1

	WM_INPUT               uintptr = 0x00FF
	WM_INPUT_DEVICE_CHANGE uintptr = 0x00FE

	GIDC_ARRIVAL uintptr = 1
	GIDC_REMOVAL uintptr = 2

	RIDEV_APPKEYS      uintptr = 0x00000400
	RIDEV_CAPTUREMOUSE uintptr = 0x00000200
	RIDEV_DEVNOTIFY    uintptr = 0x00002000
	RIDEV_EXCLUDE      uintptr = 0x00000010
	RIDEV_EXINPUTSINK  uintptr = 0x00001000
	RIDEV_INPUTSINK    uintptr = 0x00000100
	RIDEV_NOHOTKEYS    uintptr = 0x00000200
	RIDEV_NOLEGACY     uintptr = 0x00000030
	RIDEV_PAGEONLY     uintptr = 0x00000020
	RIDEV_REMOVE       uintptr = 0x00000001

	RIM_TYPEHID      uintptr = 2
	RIM_TYPEKEYBOARD uintptr = 1
	RIM_TYPEMOUSE    uintptr = 0

	KEYBOARD_OVERRUN_MAKE_CODE uintptr = 0xFF

	RI_KEY_BREAK uintptr = 1
	RI_KEY_E0    uintptr = 2
	RI_KEY_E1    uintptr = 4
	RI_KEY_MAKE  uintptr = 0

	MOUSE_ATTRIBUTES_CHANGED uintptr = 4
	MOUSE_MOVE_RELATIVE      uintptr = 0
	MOUSE_MOVE_ABSOLUTE      uintptr = 1
	MOUSE_VIRTUAL_DESKTOP    uintptr = 2

	RI_MOUSE_LEFT_BUTTON_DOWN   uintptr = 0x0001
	RI_MOUSE_LEFT_BUTTON_UP     uintptr = 0x0002
	RI_MOUSE_MIDDLE_BUTTON_DOWN uintptr = 0x0010
	RI_MOUSE_MIDDLE_BUTTON_UP   uintptr = 0x0020
	RI_MOUSE_RIGHT_BUTTON_DOWN  uintptr = 0x0004
	RI_MOUSE_RIGHT_BUTTON_UP    uintptr = 0x0008
	RI_MOUSE_BUTTON_1_DOWN      uintptr = 0x0001
	RI_MOUSE_BUTTON_1_UP        uintptr = 0x0002
	RI_MOUSE_BUTTON_2_DOWN      uintptr = 0x0004
	RI_MOUSE_BUTTON_2_UP        uintptr = 0x0008
	RI_MOUSE_BUTTON_3_DOWN      uintptr = 0x0010
	RI_MOUSE_BUTTON_3_UP        uintptr = 0x0020
	RI_MOUSE_BUTTON_4_DOWN      uintptr = 0x0040
	RI_MOUSE_BUTTON_4_UP        uintptr = 0x0080
	RI_MOUSE_BUTTON_5_DOWN      uintptr = 0x0100
	RI_MOUSE_BUTTON_5_UP        uintptr = 0x0200
	RI_MOUSE_WHEEL              uintptr = 0x0400

	RID_HEADER uintptr = 0x10000005
	RID_INPUT  uintptr = 0x10000003

	RIDI_DEVICENAME    uintptr = 0x20000007
	RIDI_DEVICEINFO    uintptr = 0x2000000B
	RIDI_PREPARSEDDATA uintptr = 0x20000005
)

const (
	_RI_MB  uintptr = 0
	_RI_MBI uintptr = 1
	_RI_K   uintptr = 2
	_RI_HID uintptr = 3
)
