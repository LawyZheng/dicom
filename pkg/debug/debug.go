package debug

import (
	"fmt"
	"log"
	"os"
)

func SetDebug(debug ...bool) {
	if len(debug) > 0 {
		_debug = debug[0]
	} else {
		_debug = true
	}
}

var _debug = false
var std = log.New(os.Stdout, "[go-dicom] ", log.Ltime|log.Lshortfile)

// Log only logs when the debug build flag is present.
func Log(data string) {
	if _debug {
		std.Output(3, data)
	}

}

// Logf only logs with a formatted string when the debug build flag is present.
func Logf(format string, args ...interface{}) {
	if _debug {
		std.Output(3, fmt.Sprintf(format, args...))
	}

}
