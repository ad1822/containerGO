package network

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/vishvananda/netlink"
)

func Veth() {
	veth := &netlink.Veth{
		LinkAttrs: netlink.LinkAttrs{Name: "host-veth"},
		PeerName:  "container-veth",
	}

	if err := netlink.LinkAdd(veth); err != nil {
		log.Fatalf("Failed to create veth pair: %v", err)
	}

	fmt.Println("Veth pair created: host-veth <--> container-veth")
}

func Move() {
	netnsPath := "/var/run/netns/mycontainer" // Replace with actual path

	// Get the veth interface
	containerVeth, err := netlink.LinkByName("container-veth")
	if err != nil {
		log.Fatalf("Failed to find container-veth: %v", err)
	}

	// Move the container-veth interface into the container's namespace
	if err := netlink.LinkSetNsFd(containerVeth, netnsPath); err != nil {
		log.Fatalf("Failed to move container-veth: %v", err)
	}

	fmt.Println("✅ Moved container-veth to the container's network namespace")
}

func IP() {
	// Assign IP to container-veth inside the container's namespace
	cmd := exec.Command("ip", "netns", "exec", "mycontainer", "ip", "addr", "add", "192.168.1.100/24", "dev", "container-veth")
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to assign IP: %v", err)
	}

	// Bring the interface up
	cmd = exec.Command("ip", "netns", "exec", "mycontainer", "ip", "link", "set", "container-veth", "up")
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to bring container-veth up: %v", err)
	}

	fmt.Println("✅ Assigned IP 192.168.1.100 to container-veth")
}
