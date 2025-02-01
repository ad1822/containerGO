package mount

import (
	"fmt"
	"os"
	"syscall"
)

func BindMount(source, target string) {
	if err := os.MkdirAll(target, 0755); err != nil {
		fmt.Printf("Error creating target directory %s: %v\n", target, err)
		return
	}

	if err := syscall.Mount(source, target, "", syscall.MS_BIND, ""); err != nil {
		fmt.Printf("Error bind mounting %s to %s: %v\n", source, target, err)
	}
	fmt.Printf("Bind mounted %s to %s\n", source, target)
}
