package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"text/tabwriter"
)

const dateFormat = "2006-01-02 15:04 MST"

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get [SUBSTR]",
	Aliases: []string{"g"},
	Short:   "Get details of load-balancers",
	Long:    `Get details of load-balancers, optionally matching given sub-string.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			getLoadBalancers("")
		case 1:
			getLoadBalancers(args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

// getLoadBalancers
func getLoadBalancers(substr string) {
	const format = "%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "NAME", "CREATED", "TYPE", "STATUS")
	for _, l := range loadBalancers(substr) {
		fmt.Fprintf(tw, format, *l.LoadBalancerName, (*l.CreatedTime).Format(dateFormat), l.Type, (*l.State).Code)
	}
	tw.Flush()
}
