package internal_debug
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

