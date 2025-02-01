package mount

import (
	"fmt"
	"syscall"

	"containerGO/config"
)

func MountOverlayFS() {
	err := syscall.Mount("overlay", config.MergedDir, "overlay", 0,
		fmt.Sprintf("lowerdir=%s,upperdir=%s,workdir=%s", config.LowerDir, config.UpperDir, config.WorkDir))
	if err != nil {
		fmt.Printf("Error mounting OverlayFS: %v\n", err)
		return
	}
	fmt.Println("OverlayFS mounted successfully")
}
