package commands

import (
	"containerGO/internal/utils"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

func ImageExists(imagePath string) bool {
	_, err := os.Stat(imagePath)
	return !os.IsNotExist(err) // Returns true if the image exists
}

func Run(containerName, imageName string, commandArgs []string, bindMounts []string) {
	utils.Logger(color.FgHiBlue, "🐳 Running a New Container")

	srcPath := filepath.Join(utils.GetContainerBaseDir("Images"), imageName)
	if ImageExists(srcPath) {
		utils.Logger(color.FgHiBlue, "Image is already available")
	} else {
		PullImage(imageName)
	}
	destPath := filepath.Join(utils.GetContainerBaseDir("ExtractImages"), imageName)
	if ImageExists(destPath) {
		utils.Logger(color.FgHiBlue, "Image is already extracted")
	} else {
		ExtractRootFS(srcPath, destPath)
	}

	containerBaseDir := utils.GetContainerBaseDir(containerBaseDirName)
	containerPath := filepath.Join(containerBaseDir, containerName)

	if err := SetupDir(containerName); err != nil {
		utils.Logger(color.FgRed, fmt.Sprintf("❌ Error setting up container directory: %v", err))
		os.Exit(1)
	}

	utils.Logger(color.FgGreen, fmt.Sprintf("✅ Process (PID: %v) is running on the host machine.", os.Getpid()))
	logFile, err := os.OpenFile(fmt.Sprintf("%s.log", containerName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	utils.Must(err)
	defer logFile.Close()

	// stdoutPipeReader, stdoutPipeWriter := io.Pipe()
	// stderrPipeReader, stderrPipeWriter := io.Pipe()
	// multiStdoutWriter := io.MultiWriter(os.Stdout, logFile) // Write to terminal & log
	// multiStderrWriter := io.MultiWriter(os.Stderr, logFile)

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	multiErrorWriter := io.MultiWriter(os.Stderr, logFile)

	cmd := exec.Command("/proc/self/exe", append([]string{"child", containerPath, imageName}, commandArgs...)...)

	cmd.Stdin = os.Stdin
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	cmd.Stdout = multiWriter
	cmd.Stderr = multiErrorWriter
	// cmd.Stdout = stdoutPipeWriter
	// cmd.Stderr = stderrPipeWriter
	// cmd.Stdin = os.Stdin
	// go func() {
	// 	io.Copy(multiStdoutWriter, stdoutPipeReader) // Send process stdout to terminal & log
	// }()
	// go func() {
	// 	io.Copy(multiStderrWriter, stderrPipeReader) // Send process stderr to terminal & log
	// }()
	// go func() {
	// input, err := io.MultiWriter(cmd.Stdin, logFile)
	// utils.Must(err)
	// io.Copy(input, os.Stdin) // Capture user input in log
	// }()

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

	// Set BIND_MOUNTS environment variable for the child process
	if len(bindMounts) > 0 {
		cmd.Env = append(os.Environ(), "BIND_MOUNTS="+strings.Join(bindMounts, ","))
	}

	defer utils.Cleanup(containerPath)
	if err := cmd.Run(); err != nil {
		utils.Logger(color.FgRed, fmt.Sprintf("❌ Error running container: %v", err))
		os.Exit(1)
	}
}
