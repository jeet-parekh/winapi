package winapi

type (
	// https://msdn.microsoft.com/en-us/library/ms645549.aspx
	RAWHID struct {
		SizeHid uint32
		Count uint32
		RawData []uint8
	}

	// https://msdn.microsoft.com/en-us/library/ms645562.aspx
	RAWINPUT_MB struct {
		Header RAWINPUTHEADER
		Data RAWMOUSEBUTTONS
	}

	// https://msdn.microsoft.com/en-us/library/ms645562.aspx
	RAWINPUT_MBI struct {
		Header RAWINPUTHEADER
		Data RAWMOUSEBUTTONINFO
	}

	// https://msdn.microsoft.com/en-us/library/ms645562.aspx
	RAWINPUT_K struct {
		Header RAWINPUTHEADER
		Data RAWKEYBOARD
	}

	// https://msdn.microsoft.com/en-us/library/ms645562.aspx
	RAWINPUT_HID struct {
		Header RAWINPUTHEADER
		Data RAWHID
	}

	// https://msdn.microsoft.com/en-us/library/ms645565.aspx
	RAWINPUTDEVICE struct {
		UsagePage uint16
		Usage uint16
		Flags uint32
		Target uintptr
	}

	// https://msdn.microsoft.com/en-us/library/ms645568.aspx
	RAWINPUTDEVICELIST struct {
		Device uintptr
		Type uint32
	}

	// https://msdn.microsoft.com/en-us/library/ms645571.aspx
	RAWINPUTHEADER struct {
		Type uint32
		Size uint32
		Device uintptr
		Param uintptr
	}

	// https://msdn.microsoft.com/en-us/library/ms645575.aspx
	RAWKEYBOARD struct {
		MakeCode uint16
		Flags uint16
		Reserved uint16
		VKey uint16
		Message uint32
		ExtraInformation uint32
	}

	// https://msdn.microsoft.com/en-us/library/ms645578.aspx
	RAWMOUSEBUTTONS struct {
		Flags uint16
		Buttons uint32
		RawButtons uint32
		LastX int32
		LastY int32
		ExtraInformation uint32
	}

	// https://msdn.microsoft.com/en-us/library/ms645578.aspx
	RAWMOUSEBUTTONINFO struct {
		Flags uint16
		ButtonInfo struct {
			ButtonFlags uint16
			ButtonData uint16
		}
		RawButtons uint32
		LastX int32
		LastY int32
		ExtraInformation uint32
	}

	// https://msdn.microsoft.com/en-us/library/ms645581.aspx
	RIDI_M struct {
		Size uint32
		Type uint32
		Definition RID_DEVICE_INFO_MOUSE
	}

	// https://msdn.microsoft.com/en-us/library/ms645581.aspx
	RIDI_K struct {
		Size uint32
		Type uint32
		Definition RID_DEVICE_INFO_KEYBOARD
	}

	// https://msdn.microsoft.com/en-us/library/ms645581.aspx
	RIDI_HID struct {
		Size uint32
		Type uint32
		Definition RID_DEVICE_INFO_HID
	}

	// https://msdn.microsoft.com/en-us/library/ms645584.aspx
	RID_DEVICE_INFO_HID struct {
		VendorId uint32
		ProductId uint32
		VersionNumber uint32
		UsagePage uint16
		Usage uint16
	}

	// https://msdn.microsoft.com/en-us/library/ms645587.aspx
	RID_DEVICE_INFO_KEYBOARD struct {
		Type uint32
		SubType uint32
		KeyboardMode uint32
		NumberOfFunctionKeys uint32
		NumberOfIndicators uint32
		NumberOfKeysTotal uint32
	}

	// https://msdn.microsoft.com/en-us/library/ms645589.aspx
	RID_DEVICE_INFO_MOUSE struct {
		Id uint32
		NumberOfButtons uint32
		SampleRate uint32
		HasHorizontalWheel int32
	}
)
