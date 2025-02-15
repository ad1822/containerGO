package commands

import (
	"containerGO/internal/utils"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"github.com/fatih/color"
)

const (
	containerBaseDirName = "Containers"
)

func SetupDir(containerName string) error {
	containerBaseDir := utils.GetContainerBaseDir(containerBaseDirName)
	containerPath := filepath.Join(containerBaseDir, containerName)

	if err := os.MkdirAll(containerPath, 0755); err != nil {
		return fmt.Errorf("error while creating container directory: %v", err)
	}

	overlayDirs := []string{"work", "merged", "upper"}
	for _, dir := range overlayDirs {
		dirPath := filepath.Join(containerPath, dir)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return fmt.Errorf("error creating overlay directory %s: %v", dir, err)
		}
	}

	return nil
}

func Run(containerName, imageName string, commandArgs []string) {
	utils.Logger(color.FgHiBlue, "🐳 Running a New Container")

	containerBaseDir := utils.GetContainerBaseDir(containerBaseDirName)
	containerPath := filepath.Join(containerBaseDir, containerName)

	if err := SetupDir(containerName); err != nil {
		utils.Logger(color.FgRed, fmt.Sprintf("❌ Error setting up container directory: %v", err))
		os.Exit(1)
	}

	utils.Logger(color.FgGreen, fmt.Sprintf("✅ Process (PID: %v) is running on the host machine.", os.Getpid()))

	cmd := exec.Command("/proc/self/exe", append([]string{"child", containerPath, imageName}, commandArgs...)...)
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

	defer utils.Cleanup(containerPath)
	if err := cmd.Run(); err != nil {
		utils.Logger(color.FgRed, fmt.Sprintf("❌ Error running container: %v", err))
		os.Exit(1)
	}
}
