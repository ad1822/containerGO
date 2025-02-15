package utils

import (
	"fmt"

	"github.com/fatih/color"
)

func Logger(clr color.Attribute, message string) {
	c := color.New(clr).SprintFunc()
	fmt.Println(c(message))
}
