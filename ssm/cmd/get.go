package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
	"text/tabwriter"
)

const dateFormat = "2006-01-02 15:04 MST"

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get PATH",
	Aliases: []string{"g"},
	Short:   "Get details for parameters in path",
	Long:    `Get ssm details for parameters in given path. This is much faster than 'ls' command.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 1:
			getParams(args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

// get details for params in path
func getParams(path string) {
	// path must start with /, so help out user and add if missing
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	const format = "%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "NAME", "TYPE", "VERSION", "MODIFIED")

	for _, p := range paramsByPath(path) {
		fmt.Fprintf(tw, format, *p.Name, string(p.Type), p.Version, (*p.LastModifiedDate).Format(dateFormat))
	}

	tw.Flush()
}

func init() {
	rootCmd.AddCommand(getCmd)
}
