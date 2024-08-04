package globalhelpers

import "log"

const (
	DEBUG = true
)

func DebugPrintf(format string, args ...any) {
	if DEBUG {
		log.Printf(format, args...)
	}
}

func CheckAndFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
