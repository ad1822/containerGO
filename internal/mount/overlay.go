package mount

import (
	"containerGO/internal/utils"
	"fmt"
	"path/filepath"
	"syscall"

	"github.com/fatih/color"
)

func MountOverlayFS(containerPath, lowerDir string) error {
	rootfs := filepath.Join(containerPath, "merged")
	upperDir := filepath.Join(containerPath, "upper")
	workDir := filepath.Join(containerPath, "work")

	err := syscall.Mount("overlay", rootfs, "overlay", 0, fmt.Sprintf("lowerdir=%s,upperdir=%s,workdir=%s", lowerDir, upperDir, workDir))
	if err != nil {
		return fmt.Errorf("error mounting overlay filesystem: %v", err)
	}

	utils.Logger(color.FgBlue, "OverlayFs mounted Successfully")

	return nil
}
