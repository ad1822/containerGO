package main

import (
	"fmt"
	"os"
	"strconv" // Import strconv for string-to-int conversion

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

	case "resume":
		if len(os.Args) < 3 {
			fmt.Println("Error: pause requires a PID")
			os.Exit(1)
		}
		// Convert string PID to int
		pid, err := strconv.Atoi(os.Args[2]) // Convert string to int
		if err != nil {
			fmt.Println("Error: invalid PID format")
			os.Exit(1)
		}
		commands.Resume(pid)

	case "pause":
		if len(os.Args) < 3 {
			fmt.Println("Error: pause requires a PID")
			os.Exit(1)
		}
		// Convert string PID to int
		pid, err := strconv.Atoi(os.Args[2]) // Convert string to int
		if err != nil {
			fmt.Println("Error: invalid PID format")
			os.Exit(1)
		}
		commands.Pause(pid) // Pass the PID as int

	case "stop":
		if len(os.Args) < 3 {
			fmt.Println("Error: stop requires a PID")
			os.Exit(1)
		}
		// Convert string PID to int
		pid, err := strconv.Atoi(os.Args[2]) // Convert string to int
		if err != nil {
			fmt.Println("Error: invalid PID format")
			os.Exit(1)
		}
		commands.Stop(pid) // Pass the PID as int

	default:
		fmt.Println("Invalid command")
	}
}
