package windows

//Basic Windows API data types
type (
	PVOID           uintptr
	HANDLE          PVOID
	HWND            HANDLE
	WPARAM          ULONG_PTR
	LPARAM          uintptr
	DWORD           uint32
	POINT           tagPOINT
	MSG             tagMSG
	HINSTANCE       HANDLE
	HHOOK           HANDLE
	LONG_PTR        int
	ULONG_PTR       uint
	INT_PTR         int
	UINT_PTR        uint
	UINT            uint
	WCHAR           uint16
	LPWSTR          *WCHAR
	LRESULT         LONG_PTR
	HKL             HANDLE
	LPMSG           *MSG
	BOOL            int
	HWINEVENTHOOK   HANDLE
	HMODULE         uintptr
	LONG            int32
	KBDLLHOOKSTRUCT tagKBDLLHOOKSTRUCT
	LPDWRD          *DWORD
	HOOKPROC        func(int, uintptr, uintptr) LRESULT
	WINEVENTPROC    func(hWinEventHook HWINEVENTHOOK, event uint32, hwnd HWND, idObject int32, idChild int32, idEventThread uint32, dwmsEventTime uint32) uintptr
)
