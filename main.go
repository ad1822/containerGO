package main

import (
	"fmt"
	"os"
	"strconv"

	"containerGO/commands"
	"containerGO/utils"
)

func main() {
	utils.CheckDirectories()
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <command> [args...]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		if len(os.Args) < 3 {
			fmt.Println("Usage: run <command>")
			os.Exit(1)
		}
		commands.Run(os.Args[2:])

	case "child":
		if len(os.Args) < 3 {
			fmt.Println("Usage: child <command>")
			os.Exit(1)
		}
		commands.Child(os.Args[2:])

	case "pull":
		if len(os.Args) < 3 {
			fmt.Println("Usage: pull <image-name>")
			os.Exit(1)
		}
		err := commands.PullImage(os.Args[2])
		if err != nil {
			fmt.Println("Error pulling image:", err)
			os.Exit(1)
		}

	case "extract":
		if len(os.Args) < 3 {
			fmt.Println("Usage: extract <image-name>")
			os.Exit(1)
		}
		imageName := os.Args[2]
		srcPath := fmt.Sprintf("/home/arcadian/Downloads/ContainerGO/Images/%s", imageName)
		destPath := fmt.Sprintf("/home/arcadian/Downloads/ContainerGO/ExtractImages/%s", imageName)

		err := utils.ExtractRootFS(srcPath, destPath)
		if err != nil {
			fmt.Println("Error extracting rootfs:", err)
			os.Exit(1)
		}
		fmt.Println("Root filesystem extracted successfully!")

	case "resume":
		if len(os.Args) < 3 {
			fmt.Println("Error: resume requires a PID")
			os.Exit(1)
		}
		pid, err := strconv.Atoi(os.Args[2])
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
		pid, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: invalid PID format")
			os.Exit(1)
		}
		commands.Pause(pid)

	case "stop":
		if len(os.Args) < 3 {
			fmt.Println("Error: stop requires a PID")
			os.Exit(1)
		}
		pid, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: invalid PID format")
			os.Exit(1)
		}
		commands.Stop(pid)

	default:
		fmt.Println("Invalid command")
		os.Exit(1)
	}
}
