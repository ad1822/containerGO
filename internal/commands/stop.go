package commands

import (
	"fmt"
	"os"
	"syscall"
)

func Stop(pid int) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("invalid PID %d: %v", pid, err)
	}

	if err := process.Signal(syscall.SIGTERM); err != nil {
		return fmt.Errorf("error sending SIGTERM to PID %d: %v", pid, err)
	}

	if err := process.Signal(syscall.SIGKILL); err != nil {
		return fmt.Errorf("error sending SIGKILL to PID %d: %v", pid, err)
	}

	fmt.Printf("Container with PID %d stopped\n", pid)
	return nil
}
