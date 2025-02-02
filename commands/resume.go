package commands

import (
	"fmt"
	"os"
	"syscall"
)

func Resume(pid int) {
	err := syscall.Kill(pid, syscall.SIGCONT)
	if err != nil {
		fmt.Printf("Error resuming process with PID %d: %s\n", pid, err)
		os.Exit(1)
	}
	fmt.Printf("Process with PID %d resumed successfully.\n", pid)
}
