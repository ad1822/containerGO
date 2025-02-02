package commands

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"containerGO/utils"
)

func Run(args []string) {
	fmt.Printf("Running %v as pid %d\n", args, os.Getpid())

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, args...)...)
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

	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting container:", err)
		os.Exit(1)
	}

	fmt.Println("Container is now running with PID", cmd.Process.Pid)

	utils.Must(cmd.Wait())
}
