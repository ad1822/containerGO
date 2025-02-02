package commands

import (
	"fmt"
	"syscall"
)

func Stop(pid int) error {
	if err := syscall.Kill(pid, syscall.SIGKILL); err != nil {
		return fmt.Errorf("error stopping container with PID %d: %v", pid, err)
	}
	fmt.Printf("Container with PID %d stopped\n", pid)
	return nil
}
