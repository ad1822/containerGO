package cmd

import (
	"containerGO/internal/commands"

	"github.com/spf13/cobra"
)

var (
	containerName string
	bindMounts    []string
)

var runCmd = &cobra.Command{
	Use:   "run [flags] <image-name> <command>",
	Short: "Run a new container",
	Args:  cobra.MinimumNArgs(2), // Requires at least image-name and command
	Run: func(cmd *cobra.Command, args []string) {
		imageName := args[0]
		commandArgs := args[1:]

		// Call your Run function with the bind mounts
		commands.Run(containerName, imageName, commandArgs, bindMounts)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Add the --bind flag
	runCmd.Flags().StringSliceVarP(&bindMounts, "bind", "b", []string{}, "Bind mount a host directory into the container (format: /host/path:/container/path)")
	runCmd.Flags().StringVarP(&containerName, "name", "n", "", "Name of the container")
}
