package util

import "github.com/fatih/color"

var (
	colorList = []func(string, ...interface{}) string{
		color.BlueString,
		color.CyanString,
		color.GreenString,
		color.MagentaString,
		color.YellowString,
	}
	colorCounter int
)

func RandomizeColor(s string) string {
	colorCounter++
	if colorCounter == len(colorList) {
		colorCounter = 0
	}
	return colorList[colorCounter](s)
}
func ErrorColor(s string) string {
	return color.RedString(s)
}
