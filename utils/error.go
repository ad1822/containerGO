package utils

import "log"

func Must(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %v", err)
	}
}
