package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
	"text/tabwriter"
)

// valuesCmd represents the values command
var valuesCmd = &cobra.Command{
	Use:     "values PATH",
	Aliases: []string{"v"},
	Short:   "Show values for params in PATH",
	Long:    `Show values for all params in given path.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 1:
			getValues(args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

// get values for params in path
func getValues(path string) {
	// path must start with /, so help out user and add if missing
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	const format = "%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "NAME", "VALUE")

	for _, p := range paramsByPath(path) {
		fmt.Fprintf(tw, format, *p.Name, *p.Value)
	}

	tw.Flush()
}

func init() {
	rootCmd.AddCommand(valuesCmd)
}
