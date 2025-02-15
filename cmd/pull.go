package cmd

import (
	"fmt"

	"containerGO/internal/commands"

	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull <image-name>",
	Short: "Pull an image from a registry",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := commands.PullImage(args[0])
		if err != nil {
			fmt.Println("Error pulling image:", err)
			return
		}
		fmt.Println("Image pulled successfully!")
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
}
