package cmd

import (
	"fmt"
	"strconv"

	"containerGO/internal/commands"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop <PID>",
	Short: "Stop a running container",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pid, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: invalid PID format")
			return
		}
		if err := commands.Stop(pid); err != nil {
			fmt.Println("Error stopping container:", err)
			return
		}
		fmt.Println("Container stopped successfully!")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
