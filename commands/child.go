package commands

import (
	"containerGO/mount"
	"containerGO/utils"
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	"github.com/fatih/color"
)

// Child executes a new process inside the isolated environment
func Child(args []string) {
	utils.Logger(color.FgGreen, fmt.Sprintf("✅ Container process started successfully! (PID: %v)", os.Getpid()))

	if len(args) < 3 {
		fmt.Println("Usage: child <container-path> <image-name> <command>")
		os.Exit(1)
	}

	containerPath := args[0]
	imageName := args[1]
	rootfs := filepath.Join(containerPath, "merged")

	lowerDir := filepath.Join("/home/arcadian/Downloads/ContainerGO/ExtractImage", imageName)

	utils.Must(syscall.Sethostname([]byte("container")))

	err := mount.MountOverlayFS(containerPath, lowerDir)
	if err != nil {
		fmt.Println("Error mounting OverlayFS:", err)
		os.Exit(1)
	}

	if err := syscall.Chroot(rootfs); err != nil {
		fmt.Println("Error changing root:", err)
		os.Exit(1)
	}

	if err := os.Chdir("/"); err != nil {
		fmt.Println("Error changing directory:", err)
		os.Exit(1)
	}

	utils.Must(syscall.Mount("proc", "proc", "proc", 0, ""))

	command := args[2:]
	utils.Logger(color.FgCyan, fmt.Sprintf("Command: %v", command))

	// Execute command inside container
	err = syscall.Exec(command[0], command, os.Environ())
	if err != nil {
		fmt.Println("Error executing command inside container:", err)
		os.Exit(1)
	}
}
