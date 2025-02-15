package cmd

import (
	"fmt"
	"strconv"

	"containerGO/internal/commands"

	"github.com/spf13/cobra"
)

var resumeCmd = &cobra.Command{
	Use:   "resume <PID>",
	Short: "Resume a paused container",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pid, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: invalid PID format")
			return
		}
		if err := commands.Resume(pid); err != nil {
			fmt.Println("Error resuming container:", err)
			return
		}
		fmt.Println("Container resumed successfully!")
	},
}

func init() {
	rootCmd.AddCommand(resumeCmd)
}
