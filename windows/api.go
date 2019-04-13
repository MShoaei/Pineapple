package windows

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var kernel32 = windows.NewLazySystemDLL("kernel32.dll")
var user32 = windows.NewLazySystemDLL("user32.dll")

var (
	setWindowsHookExA        = user32.NewProc("SetWindowsHookExA")
	getWindowThreadProcessID = user32.NewProc("GetWindowThreadProcessId")
	toUnicodeEx              = user32.NewProc("ToUnicodeEx")
	translateMessage         = user32.NewProc("TranslateMessage")
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
)

func Test() {

	if err := setWindowsHookExA.Find(); err != nil {
		fmt.Println("did not find")
	}
	if err := getWindowThreadProcessID.Find(); err != nil {
		fmt.Println("did not find GetWindowThreadProcessId")
	}
	if err := toUnicodeEx.Find(); err != nil {
		fmt.Println("did not find ToUnicodeEx")
	}
	if err := translateMessage.Find(); err != nil {
		fmt.Println("did not find TranslateMessage")
	}
	if err := getForegroundWindow.Find(); err != nil {
		fmt.Println("did not find GetForegroundWindow")
	}
	if err := setWinEventHook.Find(); err != nil {
		fmt.Println("did not find SetWinEventHook")
	}
	if err := setWinEventHook.Find(); err != nil {
		fmt.Println("did not find SetWinEventHook")
	}
	if err := getMessageW.Find(); err != nil {
		fmt.Println("did not find GetMessage")
	}
	if err := unhookWindowsHookEx.Find(); err != nil {
		fmt.Println("did not find UnhookWindowsHookEx")
	}
	if err := mapVirtualKeyExW.Find(); err != nil {
		fmt.Println("did not find MapVirtualKeyExW")
	}
	if err := getKeyboardLayout.Find(); err != nil {
		fmt.Println("did not find GetKeyboardLayout")
	}
}

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

func CallNextHookEx(hhk HHOOK, nCode int, wParam WPARAM, lParam LPARAM) LRESULT {
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

func UnhookWindowsHookEx(hhk HHOOK) BOOL {
	r1, _, _ := unhookWindowsHookEx.Call(uintptr(hhk))
	return BOOL(r1)
}

func SetWinEventHook(eventMin DWORD, eventMax DWORD, hmodWinEventProc HMODULE, pfnWinEventProc WINEVENTPROC, idProcess DWORD, idThread DWORD, dwFlags DWORD) HWINEVENTHOOK {
	r1, _, _ := setWinEventHook.Call(
		uintptr(eventMin),
		uintptr(eventMax),
		uintptr(hmodWinEventProc),
		syscall.NewCallback(pfnWinEventProc),
		uintptr(idProcess),
		uintptr(idThread),
		uintptr(dwFlags))
	return HWINEVENTHOOK(r1)
}

func GetWindowTextLengthW(hwnd HWND) int {
	r1, _, _ := getWindowTextLengthW.Call(
		uintptr(hwnd))

	return int(r1)
}

func GetWindowTextW(hwnd HWND) string {
	textLen := GetWindowTextLengthW(hwnd) + 1

	buf := make([]uint16, textLen)
	getWindowTextW.Call(
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
