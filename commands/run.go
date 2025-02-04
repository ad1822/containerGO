package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"containerGO/utils"
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
	overlayDirs := []string{"work", "merged", "base", "upper"}
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

	// Create the child process with the image name passed in the arguments
	cmd := exec.Command("/proc/self/exe", append([]string{"child", containerPath, imageName}, filteredArgs...)...)
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
	fmt.Println("RUN")
	utils.Must(cmd.Run())
}
