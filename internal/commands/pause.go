package commands

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"syscall"
)

func getChildPID(parentPID int) (int, error) {
	path := fmt.Sprintf("/proc/%d/task/%d/children", parentPID, parentPID)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, fmt.Errorf("failed to read children of PID %d: %v", parentPID, err)
	}

	children := strings.Fields(string(data))
	if len(children) == 0 {
		return 0, fmt.Errorf("no child process found for PID %d", parentPID)
	}

	return strconv.Atoi(children[0])
}

func Pause(pid int) error {
	childPID, err := getChildPID(pid)
	if err != nil {
		return err
	}

	if err := syscall.Kill(childPID, syscall.SIGSTOP); err != nil {
		return fmt.Errorf("error pausing container with PID %d (bash PID: %d): %v", pid, childPID, err)
	}

	fmt.Printf("Container process (PID %d) paused\n", childPID)
	return nil
}
