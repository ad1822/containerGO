package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"containerGO/utils"

	"github.com/fatih/color"
)

const containerBaseDir = "/home/arcadian/Downloads/ContainerGO/Containers/"

func SetUpDir(containerName string) {

	containerPath := filepath.Join(containerBaseDir, containerName)
	err := os.MkdirAll(containerPath, 0755)
	if err != nil {
		fmt.Println("Error creating container directory:", err)
		os.Exit(1)
	}

	// Setup OverlayFS directories
	overlayDirs := []string{"work", "merged", "upper"}
	for _, dir := range overlayDirs {
		err := os.MkdirAll(filepath.Join(containerPath, dir), 0755)
		if err != nil {
			fmt.Println("Error creating overlay directory:", dir, err)
			os.Exit(1)
		}
	}
}

func Run(args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: run --name <container-name> <image-name> <command>")
		os.Exit(1)
	}

	utils.Logger(color.FgHiBlue, "🐳 Running a New Container")

	var containerName, imageName string
	var filteredArgs []string

	for i := 0; i < len(args); i++ {
		if args[i] == "--name" {
			if i+1 >= len(args) {
				fmt.Println("Error: --name flag requires a container name")
				os.Exit(1)
			}
			containerName = args[i+1]
			i++
		} else if len(args[i]) > 0 && args[i][0] != '-' && imageName == "" {
			imageName = args[i]
		} else {
			filteredArgs = append(filteredArgs, args[i])
		}
	}

	if containerName == "" {
		fmt.Println("Error: You must specify a container name using --name <container-name>")
		os.Exit(1)
	}

	if imageName == "" {
		fmt.Println("Error: You must specify an image name")
		os.Exit(1)
	}

	containerPath := filepath.Join(containerBaseDir, containerName)

	SetUpDir(containerName)

	utils.Logger(color.FgGreen, fmt.Sprintf("✅ Process (PID: %v) is running on the host machine.", os.Getpid()))

	cmd := exec.Command("/proc/self/exe", append([]string{"child", containerPath, imageName}, filteredArgs...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWUSER | syscall.CLONE_NEWNET,
		Unshareflags: syscall.CLONE_NEWNS,
		Credential:   &syscall.Credential{Uid: 0, Gid: 0},
		UidMappings: []syscall.SysProcIDMap{
			{ContainerID: 0, HostID: os.Geteuid(), Size: 1},
		},
		GidMappings: []syscall.SysProcIDMap{
			{ContainerID: 0, HostID: os.Getegid(), Size: 1},
		},
	}

	utils.Must(cmd.Run())
	// pid := cmd.Process.Pid
	// fmt.Println("Started process with PID:", pid)

	// if err := utils.CreateCgroup("my_container", pid); err != nil {
	// fmt.Println("Failed to create cgroup:", err)
	// os.Exit(1)

}
