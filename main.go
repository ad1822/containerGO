package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

const (
	lowerDir  = "/home/arcadian/containerfs/layers/base"
	upperDir  = "/home/arcadian/containerfs/layers/container"
	workDir   = "/home/arcadian/containerfs/layers/work"
	mergedDir = "/home/arcadian/containerfs/layers/merged"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run ./container.go run <command>")
		return
	}

	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("Invalid command")
	}
}

func run() {
	fmt.Printf("Running %v as pid %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER | syscall.CLONE_NEWNET,
		Credential: &syscall.Credential{Uid: 0, Gid: 0},
		UidMappings: []syscall.SysProcIDMap{
			{ContainerID: 0, HostID: os.Geteuid(), Size: 1},
		},
		GidMappings: []syscall.SysProcIDMap{
			{ContainerID: 0, HostID: os.Getegid(), Size: 1},
		},
	}

	must(cmd.Run())
}

func child() {
	fmt.Printf("Running %v as pid %d\n", os.Args[2:], os.Getpid())

	// Debugging: Check if directories exist
	checkDirectories()

	// Mount OverlayFS
	err := syscall.Mount("overlay", mergedDir, "overlay", 0,
		fmt.Sprintf("lowerdir=%s,upperdir=%s,workdir=%s", lowerDir, upperDir, workDir))
	if err != nil {
		fmt.Printf("Error mounting OverlayFS: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("OverlayFS mounted successfully")

	hostDir := "/home/arcadian/DEMO"    // Directory on the host
	containerDir := mergedDir + "/DEMO" // Path inside the container
	bindMount(hostDir, containerDir)

	// Set hostname
	must(syscall.Sethostname([]byte("container")))

	// Chroot into merged dir
	must(syscall.Chroot("/home/arcadian/containerfs/layers/merged"))
	must(syscall.Chdir("/"))

	// Mount /proc inside the container
	must(syscall.Mount("proc", "proc", "proc", 0, ""))

	// Cleanup function to ensure that everything is cleaned up
	// defer cleanup()

	// Execute the command inside the container
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	must(cmd.Run())
}

func bindMount(source, target string) {
	if err := os.MkdirAll(target, 0755); err != nil {
		fmt.Printf("Error creating target directory %s: %v\n", target, err)
		os.Exit(1)
	}

	if err := syscall.Mount(source, target, "", syscall.MS_BIND, ""); err != nil {
		fmt.Printf("Error bind mounting %s to %s: %v\n", source, target, err)
		os.Exit(1)
	}

	fmt.Printf("Bind mounted %s to %s\n", source, target)
}

// func cleanup() {
// 	// Unmount /proc and OverlayFS after the process terminates
// 	fmt.Println("Cleaning up...")

// 	// Unmount /proc inside the container
// 	if err := syscall.Unmount("/proc", 0); err != nil {
// 		fmt.Printf("Warning: Failed to unmount /proc: %v\n", err)
// 	} else {
// 		fmt.Println("Unmounted /proc successfully.")
// 	}

// 	// Unmount OverlayFS
// 	if err := syscall.Unmount("/", 0); err != nil {
// 		fmt.Printf("Warning: Failed to unmount /: %v\n", err)
// 	} else {
// 		fmt.Println("Unmounted OverlayFS successfully.")
// 	}

// 	// Optional: Remove temporary directories or files if needed
// 	// Cleanup any other directories you used
// 	removeContentsIfExists(upperDir)
// 	removeContentsIfExists(workDir)
// }

func removeContentsIfExists(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Printf("Warning: %s does not exist, skipping cleanup.\n", dir)
		return
	}
	removeContents(dir)
}

func removeContents(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error reading directory %s: %v\n", dir, err)
		return
	}
	for _, file := range files {
		err := os.RemoveAll(dir + "/" + file.Name())
		if err != nil {
			fmt.Printf("Error removing %s: %v\n", file.Name(), err)
		}
	}
}

// Check if a directory is mounted
func isMounted(mountPoint string) bool {
	output, err := os.ReadFile("/proc/mounts")
	if err != nil {
		fmt.Printf("Failed to read /proc/mounts: %v\n", err)
		return false
	}
	return strings.Contains(string(output), mountPoint)
}

// Debugging function to print current mounts
func printMounts() {
	fmt.Println("Current Mounts:")
	output, err := os.ReadFile("/proc/mounts")
	if err != nil {
		fmt.Printf("Failed to read /proc/mounts: %v\n", err)
		return
	}
	fmt.Println(string(output))
}

// Debugging function to check if required directories exist
func checkDirectories() {
	fmt.Println("Checking required directories...")
	for _, dir := range []string{lowerDir, upperDir, workDir, mergedDir} {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			fmt.Printf("Warning: Directory %s does not exist!\n", dir)
		}
	}
	fmt.Println("Directory check complete.")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
