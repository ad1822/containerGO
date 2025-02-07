package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// Define cgroup base path (for cgroup v2)
const cgroupBase = "/sys/fs/cgroup/my_container_runtime"

func CreateCgroup(cgroupName string, pid int) error {
	// Define the cgroup path
	cgroupPath := filepath.Join(cgroupBase, cgroupName)

	// Create the cgroup directory
	if err := os.MkdirAll(cgroupPath, 0755); err != nil {
		return fmt.Errorf("failed to create cgroup: %v", err)
	}

	// Set CPU limit (50% of a single core)
	if err := os.WriteFile(filepath.Join(cgroupPath, "cpu.max"), []byte("50000 100000"), 0644); err != nil {
		return fmt.Errorf("failed to set CPU limit: %v", err)
	}

	// Set Memory limit (512MB)
	if err := os.WriteFile(filepath.Join(cgroupPath, "memory.max"), []byte(strconv.Itoa(512*1024*1024)), 0644); err != nil {
		return fmt.Errorf("failed to set Memory limit: %v", err)
	}

	// Set Max Process Count (10 processes)
	if err := os.WriteFile(filepath.Join(cgroupPath, "pids.max"), []byte("10"), 0644); err != nil {
		return fmt.Errorf("failed to set PID limit: %v", err)
	}

	// Add the process to the cgroup
	if err := os.WriteFile(filepath.Join(cgroupPath, "cgroup.procs"), []byte(strconv.Itoa(pid)), 0644); err != nil {
		return fmt.Errorf("failed to add process to cgroup: %v", err)
	}

	fmt.Printf("Cgroup %s created and process %d added\n", cgroupName, pid)
	return nil
}
