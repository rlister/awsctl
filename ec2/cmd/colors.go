package cmd

import (
	"github.com/fatih/color"
)

var green = color.New(color.FgGreen).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()

// statusColor returns stack status as a colored string
func statusColor(s string) string {
	switch {
	case s == "available":
		return green(s)
	case s == "running":
		return green(s)
	case s == "terminated":
		return red(s)
	default:
		return yellow(s)
	}
}
