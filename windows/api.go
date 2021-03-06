package windows

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var kernel32 = windows.NewLazySystemDLL("kernel32.dll")
var user32 = windows.NewLazySystemDLL("user32.dll")

var (
	procSetFileAttributesW   = kernel32.NewProc("SetFileAttributesW")
	setWindowsHookExA        = user32.NewProc("SetWindowsHookExA")
	getWindowThreadProcessID = user32.NewProc("GetWindowThreadProcessId")
	toUnicodeEx              = user32.NewProc("ToUnicodeEx")
	getForegroundWindow      = user32.NewProc("GetForegroundWindow")
	setWinEventHook          = user32.NewProc("SetWinEventHook")
	callNextHookEx           = user32.NewProc("CallNextHookEx")
	getMessageW              = user32.NewProc("GetMessageW")
	unhookWindowsHookEx      = user32.NewProc("UnhookWindowsHookEx")
	getWindowTextLengthW     = user32.NewProc("GetWindowTextLengthW")
	getWindowTextW           = user32.NewProc("GetWindowTextW")
	mapVirtualKeyExW         = user32.NewProc("MapVirtualKeyExW")
	getKeyboardLayout        = user32.NewProc("GetKeyboardLayout")
	getKeyboardState         = user32.NewProc("GetKeyboardState")
	unhookWinEvent           = user32.NewProc("UnhookWinEvent")
	getAsyncKeyState         = user32.NewProc("GetAsyncKeyState")
	getKeyState              = user32.NewProc("GetKeyState")
	//translateMessage         = user32.NewProc("TranslateMessage")

)

func SetWindowsHookExA(idHook int, lpfn HOOKPROC, hmod HINSTANCE, dwThreadID DWORD) HHOOK {
	r1, _, _ := setWindowsHookExA.Call(
		uintptr(idHook),
		windows.NewCallback(lpfn),
		uintptr(hmod),
		uintptr(dwThreadID),
	)
	return HHOOK(r1)
	// return r1
}

func CallNextHookEx(hhk HHOOK, nCode int, wParam, lParam uintptr) LRESULT {
	r1, _, _ := callNextHookEx.Call(
		uintptr(hhk),
		uintptr(nCode),
		uintptr(wParam),
		uintptr(lParam),
	)
	return LRESULT(r1)
}

func GetMessageW(lpMSG LPMSG, hwnd HWND, wMsgFilterMin UINT, wMsgFilterMax UINT) BOOL {
	r1, _, _ := getMessageW.Call(
		uintptr(unsafe.Pointer(lpMSG)),
		uintptr(hwnd),
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax),
	)
	return BOOL(r1)
}

func UnhookWindowsHookEx(hhk HHOOK) bool {
	ret, _, _ := unhookWindowsHookEx.Call(uintptr(hhk))
	return ret != 0
}

func SetWinEventHook(eventMin uint32, eventMax uint32, hmodWinEventProc HMODULE, callbackFunction WINEVENTPROC, idProcess uint32, idThread uint32, dwFlags uint32) (HWINEVENTHOOK, error) {
	ret, _, err := syscall.Syscall9(setWinEventHook.Addr(), 7,
		uintptr(eventMin),
		uintptr(eventMax),
		uintptr(hmodWinEventProc),
		windows.NewCallback(callbackFunction),
		uintptr(idProcess),
		uintptr(idThread),
		uintptr(dwFlags),
		0, 0)

	if ret == 0 {
		return 0, err
	}

	return HWINEVENTHOOK(ret), nil
}

func GetWindowTextLengthW(hwnd HWND) int {
	r1, _, _ := getWindowTextLengthW.Call(
		uintptr(hwnd))

	return int(r1)
}

func GetWindowTextW(hwnd HWND) string {
	textLen := GetWindowTextLengthW(hwnd) + 1

	buf := make([]uint16, textLen)
	_, _, _ = getWindowTextW.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(textLen))

	return syscall.UTF16ToString(buf)
}

func MapVirtualKeyExW(uCode UINT, uMapType UINT, hkl HKL) UINT {
	r1, _, _ := mapVirtualKeyExW.Call(
		uintptr(uCode),
		uintptr(uMapType),
		uintptr(hkl),
	)
	return UINT(r1)
}

func GetKeyboardLayout(idThread DWORD) HKL {
	r1, _, _ := getKeyboardLayout.Call(
		uintptr(idThread),
	)
	return HKL(r1)
}

func ToUnicodeEx(wVirtualKey UINT, wScanCode UINT, lpKeyState *[]byte, pwszBuff *[]rune, cchBuff int, wFlags UINT, dwhkl HKL) int {
	r1, _, _ := toUnicodeEx.Call(
		uintptr(wVirtualKey),
		uintptr(wScanCode),
		uintptr(unsafe.Pointer(&(*lpKeyState)[0])),
		uintptr(unsafe.Pointer(&(*pwszBuff)[0])),
		uintptr(cchBuff),
		uintptr(wFlags),
		uintptr(dwhkl),
	)
	return int(r1)
}

func GetWindowThreadProcessID(hwnd HWND, lpdwProcessId LPDWRD) HANDLE {
	ret, _, _ := getWindowThreadProcessID.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&lpdwProcessId)))

	return HANDLE(ret)
}

func GetForegroundWindow() HWND {
	r1, _, _ := getForegroundWindow.Call()
	return HWND(r1)
}

func GetKeyboardState(lpKeyState *[]byte) bool {
	r1, _, _ := getKeyboardState.Call(
		uintptr(unsafe.Pointer(&(*lpKeyState)[0])))
	return r1 != 0
}

func UnhookWinEvent(hWinEventHook HWINEVENTHOOK) BOOL {
	r1, _, _ := unhookWinEvent.Call(
		uintptr(hWinEventHook),
	)
	return BOOL(r1)
}

func GetAsyncKeyState(vKey int) uint16 {
	r1, _, _ := getAsyncKeyState.Call(uintptr(vKey))
	return uint16(r1)
}

func GetKeyState(nVirtKey int) uint16 {
	r1, _, _ := getKeyState.Call(
		uintptr(nVirtKey),
	)

	return uint16(r1)
}

func SetFileAttributes(name *uint16, attrs uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procSetFileAttributesW.Addr(), 2, uintptr(unsafe.Pointer(name)), uintptr(attrs), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}
