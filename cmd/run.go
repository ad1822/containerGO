package cmd

import (
	"fmt"

	"containerGO/internal/commands"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [flags] <image-name> <command>",
	Short: "Run a command in a new container",
	Args:  cobra.MinimumNArgs(2), // Ensure at least 2 arguments are provided
	Run: func(cmd *cobra.Command, args []string) {
		containerName, _ := cmd.Flags().GetString("name")
		if containerName == "" {
			fmt.Println("Error: --name flag is required")
			return
		}

		// Pass the container name, image name, and command arguments to the Run function
		commands.Run(containerName, args[0], args[1:])
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringP("name", "n", "", "Name of the container")
}
