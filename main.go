package main

import (
	"fmt"
	"os"

	"containerGO/commands"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go run [--bind /host/path:/container/path] <command>")
		return
	}

	switch os.Args[1] {
	case "run":
		commands.Run(os.Args[2:])
	case "child":
		commands.Child(os.Args[2:])
	default:
		fmt.Println("Invalid command")
	}
}
