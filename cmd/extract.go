package cmd

import (
	"fmt"
	"path/filepath"

	"containerGO/internal/commands"
	"containerGO/internal/utils"

	"github.com/spf13/cobra"
)

var extractCmd = &cobra.Command{
	Use:   "extract <image-name>",
	Short: "Extract an image from a local source",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		imageName := args[0]
		srcPath := filepath.Join(utils.GetContainerBaseDir("Images"), imageName)
		destPath := filepath.Join(utils.GetContainerBaseDir("ExtractImages"), imageName)

		err := commands.ExtractRootFS(srcPath, destPath)
		if err != nil {
			fmt.Println("Error extracting image:", err)
			return
		}
		fmt.Println("Image extracted successfully!")
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)
}
