package windows

type (
	PVOID           uintptr
	HANDLE          PVOID
	HWND            HANDLE
	WPARAM          ULONG_PTR
	LPARAM          uintptr
	DWORD           uint32
	POINT           tagPOINT
	MSG             tagMSG
	HOOKPROC        func(int, WPARAM, uintptr) LRESULT
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
	WINEVENTPROC    func(HWINEVENTHOOK, DWORD, HWND, LONG, LONG, DWORD, DWORD) uintptr
	HMODULE         HINSTANCE
	LONG            int32
	KBDLLHOOKSTRUCT tagKBDLLHOOKSTRUCT
	LPDWRD          *DWORD
)
