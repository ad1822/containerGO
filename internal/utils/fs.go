package utils

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func CheckDirectories() {
	// fmt.Println("Checking required directories...")

	dirs := []string{
		GetContainerBaseDir("Containers"),
		GetContainerBaseDir("Images"),
		GetContainerBaseDir("ExtractImages"),
	}

	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			fmt.Printf("Directory %s does not exist! Creating it...\n", dir)
			if err := os.MkdirAll(dir, 0755); err != nil {
				Logger(color.BgHiBlack, fmt.Sprintf("Error creating directory %s: %v\n", dir, err))

			} else {
				Logger(color.FgHiBlue, fmt.Sprintf("Directory %s created successfully.\n", dir))
			}
		}
	}

	// fmt.Println("Directory check complete.")
}
