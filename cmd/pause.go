package cmd

import (
	"fmt"
	"os"
	"strconv"

	"containerGO/internal/commands"

	"github.com/spf13/cobra"
)

var pauseCmd = &cobra.Command{
	Use:   "pause <PID>",
	Short: "pause a running container",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pid, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: invalid PID format")
			os.Exit(1)
		}
		if err := commands.Pause(pid); err != nil {
			fmt.Println("Error resuming container:", err)
			return
		}
		fmt.Println("Container paused successfully!")
	},
}

func init() {
	rootCmd.AddCommand(pauseCmd)
}
