package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls [SUBSTR]",
	Short: "List load-balancers",
	Long:  `List all load-balancers, optionally matching given sub-string.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			listLoadBalancers("")
		case 1:
			listLoadBalancers(args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}

// listLoadBalancers
func listLoadBalancers(substr string) {
	for _, l := range loadBalancers(substr) {
		fmt.Println(*l.LoadBalancerName)
	}
}
