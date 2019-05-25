package windows

type tagPOINT struct {
	X int32
	Y int32
}

type tagMSG struct {
	Hwnd     HWND
	Message  UINT
	WParam   WPARAM
	LParam   LPARAM
	Time     DWORD
	Pt       POINT
	LPrivate DWORD
}

type tagCWPSTRUCT struct {
	LParam  LPARAM
	WParam  WPARAM
	Message UINT
	Hwnd    HWND
}

type tagKBDLLHOOKSTRUCT struct {
	VkCode      DWORD
	ScanCode    DWORD
	Flags       DWORD
	Time        DWORD
	DwExtraInfo ULONG_PTR
}
