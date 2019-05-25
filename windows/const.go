package windows

import "syscall"

const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

const (
	EVENT_MIN               uint32 = 0x00000001
	EVENT_MAX               uint32 = 0x7FFFFFFF
	EVENT_SYSTEM_FOREGROUND uint32 = 0x0003
	WINEVENT_SKIPOWNPROCESS uint32 = 2
)

const (
	// Virtual-Key Codes
	VKBACK      = 0x08
	VKTAB       = 0x09
	VKCLEAR     = 0x0C
	VKRETURN    = 0x0D
	VKSHIFT     = 0x10
	VKCONTROL   = 0x11
	VKMENU      = 0x12
	VKPAUSE     = 0x13
	VKCAPITAL   = 0x14
	VKESCAPE    = 0x1B
	VKSPACE     = 0x20
	VKPRIOR     = 0x21
	VKNEXT      = 0x22
	VKEND       = 0x23
	VKHOME      = 0x24
	VKLEFT      = 0x25
	VKUP        = 0x26
	VKRIGHT     = 0x27
	VKDOWN      = 0x28
	VKSELECT    = 0x29
	VKPRINT     = 0x2A
	VKEXECUTE   = 0x2B
	VKSNAPSHOT  = 0x2C
	VKINSERT    = 0x2D
	VKDELETE    = 0x2E
	VKLWIN      = 0x5B
	VKRWIN      = 0x5C
	VKAPPS      = 0x5D
	VKSLEEP     = 0x5F
	VKNUMPAD0   = 0x60
	VKNUMPAD1   = 0x61
	VKNUMPAD2   = 0x62
	VKNUMPAD3   = 0x63
	VKNUMPAD4   = 0x64
	VKNUMPAD5   = 0x65
	VKNUMPAD6   = 0x66
	VKNUMPAD7   = 0x67
	VKNUMPAD8   = 0x68
	VKNUMPAD9   = 0x69
	VKMULTIPLY  = 0x6A
	VKADD       = 0x6B
	VKSEPARATOR = 0x6C
	VKSUBTRACT  = 0x6D
	VKDECIMAL   = 0x6E
	VKDIVIDE    = 0x6F
	VKF1        = 0x70
	VKF2        = 0x71
	VKF3        = 0x72
	VKF4        = 0x73
	VKF5        = 0x74
	VKF6        = 0x75
	VKF7        = 0x76
	VKF8        = 0x77
	VKF9        = 0x78
	VKF10       = 0x79
	VKF11       = 0x7A
	VKF12       = 0x7B
	VKNUMLOCK   = 0x90
	VKSCROLL    = 0x91
	VKLSHIFT    = 0xA0
	VKRSHIFT    = 0xA1
	VKLCONTROL  = 0xA2
	VKRCONTROL  = 0xA3
	VKLMENU     = 0xA4
	VKRMENU     = 0xA5
	VKOEM1      = 0xBA
	VKOEMPLUS   = 0xBB
	VKOEMCOMMA  = 0xBC
	VKOEMMINUS  = 0xBD
	VKOEMPERIOD = 0xBE
	VKOEM2      = 0xBF
	VKOEM3      = 0xC0
	VKOEM4      = 0xDB
	VKOEM5      = 0xDC
	VKOEM6      = 0xDD
	VKOEM7      = 0xDE
	VKOEM8      = 0xDF
)

const (
	FILE_ATTRIBUTE_READONLY              = 0x00000001
	FILE_ATTRIBUTE_HIDDEN                = 0x00000002
	FILE_ATTRIBUTE_SYSTEM                = 0x00000004
	FILE_ATTRIBUTE_DIRECTORY             = 0x00000010
	FILE_ATTRIBUTE_ARCHIVE               = 0x00000020
	FILE_ATTRIBUTE_DEVICE                = 0x00000040
	FILE_ATTRIBUTE_NORMAL                = 0x00000080
	FILE_ATTRIBUTE_TEMPORARY             = 0x00000100
	FILE_ATTRIBUTE_SPARSE_FILE           = 0x00000200
	FILE_ATTRIBUTE_REPARSE_POINT         = 0x00000400
	FILE_ATTRIBUTE_COMPRESSED            = 0x00000800
	FILE_ATTRIBUTE_OFFLINE               = 0x00001000
	FILE_ATTRIBUTE_NOT_CONTENT_INDEXED   = 0x00002000
	FILE_ATTRIBUTE_ENCRYPTED             = 0x00004000
	FILE_ATTRIBUTE_INTEGRITY_STREAM      = 0x00008000
	FILE_ATTRIBUTE_VIRTUAL               = 0x00010000
	FILE_ATTRIBUTE_NO_SCRUB_DATA         = 0x00020000
	FILE_ATTRIBUTE_RECALL_ON_OPEN        = 0x00040000
	FILE_ATTRIBUTE_RECALL_ON_DATA_ACCESS = 0x00400000
)
