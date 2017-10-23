package winapi

type (
	// https://msdn.microsoft.com/en-us/library/ms724197.aspx
	ANIMATIONINFO struct {
		CbSize      uint32
		IMinAnimate int32
	}

	// https://msdn.microsoft.com/en-us/library/aa379651.aspx
	AUDIODESCRIPTION struct {
		CbSize  uint32
		Enabled int32
		Locale  uint32
	}

	// https://msdn.microsoft.com/en-us/library/ms724500.aspx
	MINIMIZEDMETRICS struct {
		CbSize   uint32
		IWidth   int32
		IHorzGap int32
		IVertGap int32
		IArrange int32
	}

	// https://msdn.microsoft.com/en-us/library/ff729175.aspx
	NONCLIENTMETRICS_L struct {
		CbSize           uint32
		IBorderWidth     int32
		IScrollWidth     int32
		IScrollHeight    int32
		ICaptionWidth    int32
		ICaptionHeight   int32
		LfCaptionFont    LOGFONT
		ISmCaptionWidth  int32
		ISmCaptionHeight int32
		LfSmCaptionFont  LOGFONT
		IMenuWidth       int32
		IMenuHeight      int32
		LfMenuFont       LOGFONT
		LfStatusFont     LOGFONT
		LfMessageFont    LOGFONT
	}

	// https://msdn.microsoft.com/en-us/library/ff729175.aspx
	NONCLIENTMETRICS struct {
		CbSize             uint32
		IBorderWidth       int32
		IScrollWidth       int32
		IScrollHeight      int32
		ICaptionWidth      int32
		ICaptionHeight     int32
		LfCaptionFont      LOGFONT
		ISmCaptionWidth    int32
		ISmCaptionHeight   int32
		LfSmCaptionFont    LOGFONT
		IMenuWidth         int32
		IMenuHeight        int32
		LfMenuFont         LOGFONT
		LfStatusFont       LOGFONT
		LfMessageFont      LOGFONT
		IPaddedBorderWidth int32
	}

	// https://msdn.microsoft.com/en-us/library/dd145037.aspx
	LOGFONT struct {
		LfHeight         int32
		LfWidth          int32
		LfEscapement     int32
		LfOrientation    int32
		LfWeight         int32
		LfItalic         uint8
		LfUnderline      uint8
		LfStrikeOut      uint8
		LfCharSet        uint8
		LfOutPrecision   uint8
		LfClipPrecision  uint8
		LfQuality        uint8
		LfPitchAndFamily uint8
		LfFaceName       []uint16
	}

	// https://msdn.microsoft.com/en-us/library/ms645549.aspx
	RAWHID struct {
		DwSizeHid uint32
		DwCount   uint32
		BRawData  []uint8
	}

	// https://msdn.microsoft.com/en-us/library/ms645562.aspx
	RAWINPUT_MB struct {
		Header RAWINPUTHEADER
		Mouse  RAWMOUSEBUTTONS
	}

	// https://msdn.microsoft.com/en-us/library/ms645562.aspx
	RAWINPUT_MBI struct {
		Header RAWINPUTHEADER
		Mouse  RAWMOUSEBUTTONINFO
	}

	// https://msdn.microsoft.com/en-us/library/ms645562.aspx
	RAWINPUT_K struct {
		Header   RAWINPUTHEADER
		Keyboard RAWKEYBOARD
	}

	// https://msdn.microsoft.com/en-us/library/ms645562.aspx
	RAWINPUT_HID struct {
		Header RAWINPUTHEADER
		HID    RAWHID
	}

	// https://msdn.microsoft.com/en-us/library/ms645565.aspx
	RAWINPUTDEVICE struct {
		UsUsagePage uint16
		UsUsage     uint16
		DwFlags     uint32
		HwndTarget  uintptr
	}

	// https://msdn.microsoft.com/en-us/library/ms645568.aspx
	RAWINPUTDEVICELIST struct {
		HDevice uintptr
		DwType  uint32
	}

	// https://msdn.microsoft.com/en-us/library/ms645571.aspx
	RAWINPUTHEADER struct {
		DwType  uint32
		DwSize  uint32
		HDevice uintptr
		WParam  uintptr
	}

	// https://msdn.microsoft.com/en-us/library/ms645575.aspx
	RAWKEYBOARD struct {
		MakeCode         uint16
		Flags            uint16
		Reserved         uint16
		VKey             uint16
		Message          uint32
		ExtraInformation uint32
	}

	// https://msdn.microsoft.com/en-us/library/ms645578.aspx
	RAWMOUSEBUTTONS struct {
		UsFlags            uint16
		UlButtons          uint32
		UlRawButtons       uint32
		LLastX             int32
		LLastY             int32
		UlExtraInformation uint32
	}

	// https://msdn.microsoft.com/en-us/library/ms645578.aspx
	RAWMOUSEBUTTONINFO struct {
		UsFlags    uint16
		ButtonInfo struct {
			UsButtonFlags uint16
			UsButtonData  uint16
		}
		UlRawButtons       uint32
		LLastX             int32
		LLastY             int32
		UlExtraInformation uint32
	}

	// https://msdn.microsoft.com/en-us/library/ms645581.aspx
	RIDI_M struct {
		CbSize uint32
		DwType uint32
		Mouse  RID_DEVICE_INFO_MOUSE
	}

	// https://msdn.microsoft.com/en-us/library/ms645581.aspx
	RIDI_K struct {
		CbSize   uint32
		DwType   uint32
		Keyboard RID_DEVICE_INFO_KEYBOARD
	}

	// https://msdn.microsoft.com/en-us/library/ms645581.aspx
	RIDI_HID struct {
		CbSize uint32
		DwType uint32
		HID    RID_DEVICE_INFO_HID
	}

	// https://msdn.microsoft.com/en-us/library/ms645584.aspx
	RID_DEVICE_INFO_HID struct {
		DwVendorId      uint32
		DwProductId     uint32
		DwVersionNumber uint32
		UsUsagePage     uint16
		UsUsage         uint16
	}

	// https://msdn.microsoft.com/en-us/library/ms645587.aspx
	RID_DEVICE_INFO_KEYBOARD struct {
		DwType                 uint32
		DwSubType              uint32
		DwKeyboardMode         uint32
		DwNumberOfFunctionKeys uint32
		DwNumberOfIndicators   uint32
		DwNumberOfKeysTotal    uint32
	}

	// https://msdn.microsoft.com/en-us/library/ms645589.aspx
	RID_DEVICE_INFO_MOUSE struct {
		DwId                uint32
		DwNumberOfButtons   uint32
		DwSampleRate        uint32
		FHasHorizontalWheel int32
	}

	// https://msdn.microsoft.com/en-us/library/ms644957.aspx
	BSMINFO struct {
		CbSize uint32
		Hdesk  uintptr
		Hwnd   uintptr
		Luid   LUID
	}

	// https://msdn.microsoft.com/en-us/library/ms644958.aspx
	MSG struct {
		Hwnd    uintptr
		Message uint32
		WParam  uintptr
		LParam  uintptr
		Time    uint32
		Pt      POINT
	}

	// https://msdn.microsoft.com/en-us/library/aa379261.aspx
	LUID struct {
		LowPart  uint32
		HighPart int32
	}

	// https://msdn.microsoft.com/en-us/library/dd162805.aspx
	POINT struct {
		X int32
		Y int32
	}

	// https://msdn.microsoft.com/en-us/library/ms632603.aspx
	CREATESTRUCT struct {
		LpCreateParams uintptr
		HInstance      uintptr
		HMenu          uintptr
		HwndParent     uintptr
		Cy             int32
		Cx             int32
		Y              int32
		X              int32
		Style          int32
		LpszName       []uint16
		LpszClass      []uint16
		DwExStyle      uint32
	}

	// https://msdn.microsoft.com/en-us/library/ms632602.aspx
	CLIENTCREATESTRUCT struct {
		HWindowMenu  uintptr
		IdFirstChild uint32
	}

	// https://msdn.microsoft.com/en-us/library/ms644910.aspx
	MDICREATESTRUCT struct {
		SzClass []uint16
		SzTitle []uint16
		HOwner  uintptr
		X       int32
		Y       int32
		Cx      int32
		Cy      int32
		Style   uint32
		LParam  uintptr
	}

	// https://msdn.microsoft.com/en-us/library/ms644962.aspx
	CBT_CREATEWND struct {
		Lpcs            *CREATESTRUCT
		HwndInsertAfter uintptr
	}

	// https://msdn.microsoft.com/en-us/library/ms644961.aspx
	CBTACTIVATESTRUCT struct {
		FMouse     int32
		HwndActive uintptr
	}

	// https://msdn.microsoft.com/en-us/library/ms644963.aspx
	CWPRETSTRUCT struct {
		LResult uintptr
		LParam  uintptr
		WParam  uintptr
		Message uint32
		Hwnd    uintptr
	}

	// https://msdn.microsoft.com/en-us/library/ms644964.aspx
	CWPSTRUCT struct {
		LParam  uintptr
		WParam  uintptr
		Message uint32
		Hwnd    uintptr
	}

	// https://msdn.microsoft.com/en-us/library/ms644965.aspx
	DEBUGHOOKINFO struct {
		IdThread          uint32
		IdThreadInstaller uint32
		LParam            uintptr
		WParam            uintptr
		Code              int32
	}

	// https://msdn.microsoft.com/en-us/library/ms644966.aspx
	EVENTMSG struct {
		Message uint32
		ParamL  uint32
		ParamH  uint32
		Time    uint32
		Hwnd    uintptr
	}

	// https://msdn.microsoft.com/en-us/library/ms644967.aspx
	KBDLLHOOKSTRUCT struct {
		VkCode      uint32
		ScanCode    uint32
		Flags       uint32
		Time        uint32
		DwExtraInfo uint32
	}

	// https://msdn.microsoft.com/en-us/library/ms644968.aspx
	MOUSEHOOKSTRUCT struct {
		Pt           POINT
		Hwnd         uintptr
		WHitTestCode uint32
		DwExtraInfo  uint32
	}

	// https://msdn.microsoft.com/en-us/library/ms644969.aspx
	MOUSEHOOKSTRUCTEX struct {
		MOUSEHOOKSTRUCT MOUSEHOOKSTRUCT
		MouseData       uint32
	}

	// https://msdn.microsoft.com/en-us/library/ms644970.aspx
	MSLLHOOKSTRUCT struct {
		Pt          POINT
		MouseData   uint32
		Flags       uint32
		Time        uint32
		DwExtraInfo uint32
	}

	// https://msdn.microsoft.com/en-us/library/ms646269.aspx
	HARDWAREINPUT struct {
		UMsg    uint32
		WParamL uint16
		WParamH uint16
	}

	// https://msdn.microsoft.com/en-us/library/ms646271.aspx
	KEYBDINPUT struct {
		WVk         uint16
		WScan       uint16
		DwFlags     uint32
		Time        uint32
		DwExtraInfo uint32
	}

	// https://msdn.microsoft.com/en-us/library/ms646273.aspx
	MOUSEINPUT struct {
		Dx          int32
		Dy          int32
		MouseData   uint32
		DwFlags     uint32
		Time        uint32
		DwExtraInfo uint32
	}

	// https://msdn.microsoft.com/en-us/library/ms646270.aspx
	INPUT_M struct {
		Type uint32
		MI   MOUSEINPUT
	}

	// https://msdn.microsoft.com/en-us/library/ms646270.aspx
	INPUT_K struct {
		Type uint32
		KI   KEYBDINPUT
	}

	// https://msdn.microsoft.com/en-us/library/ms646270.aspx
	INPUT_HW struct {
		Type uint32
		HI   HARDWAREINPUT
	}

	// https://msdn.microsoft.com/en-us/library/ms646272.aspx
	LASTINPUTINFO struct {
		CbSize uint32
		DwTime uint32
	}

	// https://msdn.microsoft.com/en-us/library/ms645603.aspx
	MOUSEMOVEPOINT struct {
		X           int32
		Y           int32
		Time        uint32
		DwExtraInfo uint32
	}

	// https://msdn.microsoft.com/en-us/library/ms645604.aspx
	TRACKMOUSEEVENT struct {
		CbSize      uint32
		DwFlags     uint32
		HwndTrack   uintptr
		DwHoverTime uint32
	}

	// https://docs.microsoft.com/en-us/windows/console/coord-str
	COORD struct {
		X int16
		Y int16
	}

	// https://docs.microsoft.com/en-us/windows/console/focus-event-record-str
	FOCUS_EVENT_RECORD struct {
		BSetFocus int32
	}

	// https://docs.microsoft.com/en-us/windows/console/mouse-event-record-str
	MOUSE_EVENT_RECORD struct {
		DwMousePosition   COORD
		DwButtonState     uint32
		DwControlKeyState uint32
		DwEventFlags      uint32
	}

	// https://docs.microsoft.com/en-us/windows/console/menu-event-record-str
	MENU_EVENT_RECORD struct {
		DwCommandId uint32
	}

	// https://docs.microsoft.com/en-us/windows/console/key-event-record-str
	KEY_EVENT_RECORD struct {
		BKeyDown          int32
		WRepeatCount      uint16
		WVirtualKeyCode   uint16
		WVirtualScanCode  uint16
		UChar             uint16
		DwControlKeyState uint32
	}

	// https://docs.microsoft.com/en-us/windows/console/window-buffer-size-record-str
	WINDOW_BUFFER_SIZE_RECORD struct {
		DwSize COORD
	}

	// https://docs.microsoft.com/en-us/windows/console/input-record-str
	INPUT_RECORD_KEY struct {
		EventType uint16
		Event     KEY_EVENT_RECORD
	}

	// https://docs.microsoft.com/en-us/windows/console/input-record-str
	INPUT_RECORD_MOUSE struct {
		EventType uint16
		Event     MOUSE_EVENT_RECORD
	}

	// https://docs.microsoft.com/en-us/windows/console/input-record-str
	INPUT_RECORD_WINDOW_BUFFER_SIZE struct {
		EventType uint16
		Event     WINDOW_BUFFER_SIZE_RECORD
	}

	// https://docs.microsoft.com/en-us/windows/console/input-record-str
	INPUT_RECORD_MENU struct {
		EventType uint16
		Event     MENU_EVENT_RECORD
	}

	// https://docs.microsoft.com/en-us/windows/console/input-record-str
	INPUT_RECORD_FOCUS struct {
		EventType uint16
		Event     FOCUS_EVENT_RECORD
	}
)
