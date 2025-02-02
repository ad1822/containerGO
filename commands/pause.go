package commands

import (
	"fmt"
	"syscall"
)

func Pause(pid int) error {
	if err := syscall.Kill(pid, syscall.SIGSTOP); err != nil {
		return fmt.Errorf("error pausing container with PID %d: %v", pid, err)
	}
	fmt.Printf("Container with PID %d paused\n", pid)
	return nil
}
