package mount

import (
	"containerGO/internal/utils"
	"fmt"
	"os"
	"syscall"

	"github.com/fatih/color"
)

func BindMount(source, target string, readOnly bool) error {
	if _, err := os.Stat(source); os.IsNotExist(err) {
		return fmt.Errorf("source directory does not exist: %s", source)
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return fmt.Errorf("error creating target directory %s: %v", target, err)
	}

	flags := syscall.MS_BIND
	if readOnly {
		flags |= syscall.MS_RDONLY
	}

	if err := syscall.Mount(source, target, "", uintptr(flags), ""); err != nil {
		return fmt.Errorf("error bind mounting %s to %s: %v", source, target, err)
	}

	utils.Logger(color.FgGreen, fmt.Sprintf("Bind mounted %s to %s (read-only: %v)", source, target, readOnly))
	return nil
}
