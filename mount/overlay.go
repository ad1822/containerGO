package mount

import (
	"containerGO/utils"
	"fmt"
	"path/filepath"
	"syscall"

	"github.com/fatih/color"
)

// MountOverlayFS mounts the OverlayFS for the container
func MountOverlayFS(containerPath, lowerDir string) error {
	// Define directories for OverlayFS
	rootfs := filepath.Join(containerPath, "merged")
	upperDir := filepath.Join(containerPath, "upper")
	workDir := filepath.Join(containerPath, "work")

	// Mount the OverlayFS
	err := syscall.Mount("overlay", rootfs, "overlay", 0, fmt.Sprintf("lowerdir=%s,upperdir=%s,workdir=%s", lowerDir, upperDir, workDir))
	if err != nil {
		return fmt.Errorf("error mounting overlay filesystem: %v", err)
	}

	utils.Logger(color.FgBlue, "OverlayFs mounted Successfully")

	// fmt.Println("OverlayFS mounted successfully at", rootfs)
	return nil
}
