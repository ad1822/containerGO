package commands

import (
	"fmt"
	"os"
	"syscall"
)

// Resume resumes a paused container process
func Resume(pid int) error {
	childPid, err := getChildPID(pid)
	if err != nil {
		return fmt.Errorf("error finding child process for PID %d: %v", pid, err)
	}

	if err := syscall.Kill(childPid, syscall.SIGCONT); err != nil {
		fmt.Printf("Error resuming process with PID %d: %v\n", childPid, err)
		os.Exit(1)
	}

	fmt.Printf("Process with PID %d resumed successfully.\n", childPid)
	return nil
}
