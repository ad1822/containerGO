package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

func Err(err error, message string) {
	if err != nil {
		// fmt.Println(message, " : ", err)
		Logger(color.FgRed, fmt.Sprintf(message, " : ", err))
		os.Exit(1)
	}
}

func Must(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %v", err)
	}
}
