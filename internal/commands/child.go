package commands

import (
	"containerGO/internal/mount"
	"containerGO/internal/utils"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/fatih/color"
)

func Child(args []string) {
	utils.Logger(color.FgGreen, fmt.Sprintf("✅ Container process started successfully! (PID: %v)", os.Getpid()))

	if len(args) < 3 {
		fmt.Println("Usage: child <container-path> <image-name> <command>")
		os.Exit(1)
	}

	containerPath := args[0]
	imageName := args[1]
	rootfs := filepath.Join(containerPath, "merged")

	baseDir := utils.GetContainerBaseDir("ExtractImages")
	lowerDir := filepath.Join(baseDir, imageName)

	utils.Must(syscall.Sethostname([]byte("container")))

	err := mount.MountOverlayFS(containerPath, lowerDir)
	if err != nil {
		fmt.Println("Error mounting OverlayFS:", err)
		os.Exit(1)
	}
	bindMounts := os.Getenv("BIND_MOUNTS")
	if bindMounts != "" {
		for _, bindM := range strings.Split(bindMounts, ",") {
			parts := strings.Split(bindM, ":")
			if len(parts) < 2 || len(parts) > 3 {
				utils.Logger(color.FgRed, fmt.Sprintf("Invalid bind mount format: %s", bindM))
				continue
			}

			readOnly := false
			if len(parts) == 3 && parts[2] == "ro" {
				readOnly = true
			}

			if err := mount.BindMount(parts[0], filepath.Join(rootfs, parts[1]), readOnly); err != nil {
				utils.Logger(color.FgRed, fmt.Sprintf("Failed to bind mount: %v", err))
			}
		}
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

	err = syscall.Exec(command[0], command, os.Environ())
	if err != nil {
		fmt.Println("Error executing command inside container:", err)
		os.Exit(1)
	}
}
