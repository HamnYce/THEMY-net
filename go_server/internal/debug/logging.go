package debug

import "log"

var (
	DEBUG = false
	SEED  = false
)

func DebugPrintf(format string, args ...any) {
	if DEBUG {
		log.Printf(format, args...)
	}
}

func CheckAndFatal(err error) {
	if err == nil {
    return
	}

  log.Fatal(err)
}
