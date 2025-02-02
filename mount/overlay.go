package mount

import (
	"containerGO/config"
	"fmt"
	"os"
	"syscall"
)

// func MountOverlayFS() {
// 	err := syscall.Mount("overlay", config.MergedDir, "overlay", 0,
// 		fmt.Sprintf("lowerdir=%s,upperdir=%s,workdir=%s", config.LowerDir, config.UpperDir, config.WorkDir))
// 	if err != nil {
// 		fmt.Printf("Error mounting OverlayFS: %v\n", err)
// 		return
// 	}
// 	fmt.Println("OverlayFS mounted successfully")
// }

func MountOverlayFS() {
	// Define directories
	// lowerDir := "/tmp/container-rootfs/" // Extracted rootfs (read-only)
	// upperDir := "/tmp/container-overlay/upper"
	// workDir := "/tmp/container-overlay/work"
	// mergedDir := "/tmp/container-overlay/merged"

	// Create necessary directories
	dirs := []string{config.LowerDir, config.UpperDir, config.WorkDir}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}

	// Prepare OverlayFS options
	options := fmt.Sprintf("lowerdir=%s,upperdir=%s,workdir=%s", config.LowerDir, config.UpperDir, config.WorkDir)

	// Mount OverlayFS
	if err := syscall.Mount("overlay", config.MergedDir, "overlay", 0, options); err != nil {
		fmt.Println("Error mounting OverlayFS:", err)
		return
	}

	fmt.Println("OverlayFS mounted successfully at", config.MergedDir)
}
