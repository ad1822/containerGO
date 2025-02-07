package utils

import (
	"fmt"
	"os"
)

// CheckDirectories ensures required directories exist, creating them if necessary.
func CheckDirectories() {
	// fmt.Println("Checking required directories...")

	dirs := []string{
		"/home/arcadian/Downloads/ContainerGO/Containers/",
		"/home/arcadian/Downloads/ContainerGO/Images/",
		"/home/arcadian/Downloads/ContainerGO/ExtractImages/",
	}

	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			fmt.Printf("Directory %s does not exist! Creating it...\n", dir)
			if err := os.MkdirAll(dir, 0755); err != nil {
				fmt.Printf("Error creating directory %s: %v\n", dir, err)
			} else {
				fmt.Printf("Directory %s created successfully.\n", dir)
			}
		}
	}

	// fmt.Println("Directory check complete.")
}
