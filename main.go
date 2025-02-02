package main

import (
	"fmt"
	"os"

	"containerGO/commands"
	"containerGO/utils"
)

func main() {
	// if len(os.Args) < 3 {
	// 	fmt.Println("Usage: go run main.go run [--bind /host/path:/container/path] <command>")
	// 	return
	// }

	switch os.Args[1] {
	case "run":
		commands.Run(os.Args[2:])
	case "child":
		commands.Child(os.Args[2:])
	case "pull":
		commands.PullImage("archlinux")
	case "extract":
		err := utils.ExtractRootFS("/home/arcadian/Downloads/archlinux", "/home/arcadian/Downloads/archRootfs")
		if err != nil {
			fmt.Println("Error extracting rootfs:", err)
			return
		}
		fmt.Println("Root filesystem extracted successfully!")
	default:
		fmt.Println("Invalid command")
	}
}
