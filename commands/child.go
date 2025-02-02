package commands

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"containerGO/config"
	"containerGO/mount"
	"containerGO/utils"
)

func Child(args []string) {
	fmt.Printf("Running %v as pid %d child\n", args, os.Getpid())

	// utils.CheckDirectories()

	// Mount OverlayFS
	mount.MountOverlayFS()

	// Apply bind mounts
	// bindMounts := os.Getenv("BIND_MOUNTS")
	// if bindMounts != "" {
	// 	for _, bindM := range strings.Split(bindMounts, ",") {
	// 		parts := strings.Split(bindM, ":")
	// 		if len(parts) != 2 {
	// 			fmt.Printf("Invalid bind mount format: %s\n", bindM)
	// 			continue
	// 		}
	// 		mount.BindMount(parts[0], filepath.Join(config.MergedDir, parts[1]))
	// 	}
	// }

	// Set hostname
	utils.Must(syscall.Sethostname([]byte("container")))

	// Chroot and execute command
	if err := syscall.Chroot(config.MergedDir); err != nil {
		fmt.Println("Merged Dir Error")
	}
	// utils.Must(syscall.Chroot(config.MergedDir))
	if err := syscall.Chdir("/"); err != nil {
		fmt.Println("Chdir error")
	}
	// utils.Must(syscall.Chdir("/"))
	// utils.Must(syscall.Mount("proc", "proc", "proc", 0, ""))

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	utils.Must(cmd.Run())

	// Cleanup()

	// fmt.Println("Cleaning up container ... ")

	// if err := syscall.Unmount("/proc", 0); err != nil {
	// 	fmt.Printf("Warning: Failed to unmount /proc: %v\n", err)
	// }

	// if err := syscall.Unmount(config.MergedDir, 0); err != nil {
	// 	fmt.Printf("Warning: Failed to unmount OverlayFS: %v\n", err)
	// }

	// fmt.Println("Cleanup complete.")
}
