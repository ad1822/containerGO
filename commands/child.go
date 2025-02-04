package commands

import (
	"containerGO/mount"
	"containerGO/utils"
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

// Child executes a new process inside the isolated environment
func Child(args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: child <container-path> <image-name> <command>")
		os.Exit(1)
	}

	containerPath := args[0]
	imageName := args[1] // The image name is passed as the second argument
	rootfs := filepath.Join(containerPath, "merged")

	// Construct the lowerDir path based on the image name (this can be modified as needed)
	lowerDir := filepath.Join("/home/arcadian/Downloads/ContainerGO/ExtractImage", imageName)

	utils.Must(syscall.Sethostname([]byte("container")))

	// Ensure OverlayFS is mounted before changing root
	err := mount.MountOverlayFS(containerPath, lowerDir)
	if err != nil {
		fmt.Println("Error mounting OverlayFS:", err)
		os.Exit(1)
	}

	// Change root to the new filesystem
	if err := syscall.Chroot(rootfs); err != nil {
		fmt.Println("Error changing root:", err)
		os.Exit(1)
	}

	// Change working directory to new root
	if err := os.Chdir("/"); err != nil {
		fmt.Println("Error changing directory:", err)
		os.Exit(1)
	}

	utils.Must(syscall.Mount("proc", "proc", "proc", 0, ""))

	command := args[2:]

	// Execute command inside container
	err = syscall.Exec(command[0], command, os.Environ())
	if err != nil {
		fmt.Println("Error executing command inside container:", err)
		os.Exit(1)
	}
}
