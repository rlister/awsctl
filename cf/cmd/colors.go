package cmd

import (
	"github.com/fatih/color"
	"strings"
)

var green = color.New(color.FgGreen).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()

// statusColor returns stack status as a colored string
func statusColor(s string) string {
	switch {
	case s == "ROLLBACK_COMPLETE":
		return yellow(s)
	case strings.HasSuffix(s, "_COMPLETE"):
		return green(s)
	case strings.HasSuffix(s, "_FAILED"):
		return red(s)
	default:
		return yellow(s)
	}
}
