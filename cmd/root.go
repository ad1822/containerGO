package cmd

import (
	"containerGO/internal/commands"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "containerGO",
	Short: "A simple container runtime",
}

var childCmd = &cobra.Command{
	Use:    "child <container-path> <image-name> <command>",
	Short:  "Internal command to set up the container environment",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		commands.Child(args)
	},
}

func Execute() {
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(childCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
