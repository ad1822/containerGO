package utils

import (
	"fmt"
	"os"
)

func CheckDirectories() {
	fmt.Println("Checking required directories...")
	for _, dir := range []string{"/home/arcadian/containerfs/layers/base",
		"/home/arcadian/containerfs/layers/container",
		"/home/arcadian/containerfs/layers/work",
		"/home/arcadian/containerfs/layers/merged"} {

		if _, err := os.Stat(dir); os.IsNotExist(err) {
			fmt.Printf("Warning: Directory %s does not exist!\n", dir)
		}
	}
	fmt.Println("Directory check complete.")
}
