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

	// var bindMounts []string
	// var filteredArgs []string

	// Parse arguments for bind mounts
	// for i := 0; i < len(args); i++ {
	// 	if args[i] == "--bind" {
	// 		if i+1 >= len(args) {
	// 			fmt.Println("Error: --bind requires a value in the format /host/path:/container/path")
	// 			os.Exit(1)
	// 		}
	// 		bindMounts = append(bindMounts, args[i+1])
	// 		i++ // Skip next argument
	// 		continue
	// 	}
	// 	filteredArgs = append(filteredArgs, args[i])
	// }

	// Create the child process
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

	// Pass bind mounts as environment variables
	// if len(bindMounts) > 0 {
	// 	cmd.Env = append(os.Environ(), "BIND_MOUNTS="+strings.Join(bindMounts, ","))
	// }

	utils.Must(cmd.Run())
}
