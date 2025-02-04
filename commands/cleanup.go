package commands

import (
	"fmt"
	"syscall"

	"github.com/charmbracelet/log"
)

func Cleanup() {
	log.Info("Cleaning up container ... ")

	if err := syscall.Unmount("/proc", 0); err != nil {
		fmt.Printf("Warning: Failed to unmount /proc: %v\n", err)
	}

	if err := syscall.Unmount("/", 0); err != nil {
		fmt.Printf("Warning: Failed to unmount OverlayFS: %v\n", err)
	}

	// removeDirectoryContents(config.MergedDir)
	// removeDirectoryContents(config.UpperDir)
	// removeDirectoryContents(config.LowerDir)
	// removeDirectoryContents(config.WorkDir)

	log.Info("Cleanup complete.")
}

// func removeDirectoryContents(dir string) {
// 	if _, err := os.Stat(dir); os.IsNotExist(err) {
// 		fmt.Printf("Warning: Directory %s does not exist, skipping cleanup.\n", dir)
// 		return
// 	}

// 	files, err := os.ReadDir(dir)
// 	if err != nil {
// 		fmt.Printf("Error reading directory %s: %v\n", dir, err)
// 		return
// 	}

// 	for _, file := range files {
// 		err := os.RemoveAll(strings.Join([]string{dir, file.Name()}, "/"))
// 		if err != nil {
// 			fmt.Printf("Error removing %s: %v\n", file.Name(), err)
// 		}
// 	}
// }
