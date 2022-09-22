package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

const dateFormat = "2006-01-02 15:04 MST"

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get PREFIX",
	Aliases: []string{"g"},
	Short:   "get details for stacks",
	Long:    `List details of stacks with prefix.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			getStacks("")
		default:
			getStacks(args[0])
		}
	},
}

// getStacks shows stack name, status and creation date for all stacks with prefix
func getStacks(prefix string) {
	const format = "%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "NAME", "CREATED", "STATUS")
	for _, s := range stacks(prefix) {
		fmt.Fprintf(tw, format, *s.StackName, (*s.CreationTime).Format(dateFormat), statusColor(string(s.StackStatus)))
	}
	tw.Flush()
}

func init() {
	rootCmd.AddCommand(getCmd)
}
