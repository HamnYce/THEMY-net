package internal_debug

import "log"

func CheckAndFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
