package log

import (
	"fmt"
	"github.com/fatih/color"
)

func Debug(msg string) {
	fmt.Println(color.BlackString(msg))
}

func Error(msg string) {
	fmt.Println(color.RedString(msg))
}

func Warn(msg string) {
	fmt.Println(color.BlueString(msg))
}
