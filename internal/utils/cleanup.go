package utils

import (
	"os"

	"github.com/fatih/color"
)

func Cleanup(containerName string) {
	Logger(color.FgGreen, "Cleaning up container ..")

	// if err := syscall.Unmount("/proc", 0); err != nil {
	// 	fmt.Printf("Warning: Failed to unmount /proc: %v\n", err)
	// }

	// if err := syscall.Unmount("/", 0); err != nil {
	// 	fmt.Printf("Warning: Failed to unmount OverlayFS: %v\n", err)
	// }

	// removeDirectoryContents(config.MergedDir)
	// removeDirectoryContents(config.UpperDir)
	// removeDirectoryContents(config.LowerDir)
	// removeDirectoryContents(config.WorkDir)

	os.RemoveAll(containerName)
	Logger(color.FgGreen, "Cleanup complete ...")

}
