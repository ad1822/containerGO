package main

import (
	"containerGO/cmd"
	"containerGO/internal/utils"
)

func main() {
	utils.CheckDirectories()
	cmd.Execute()
}
