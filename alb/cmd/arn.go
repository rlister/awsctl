package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// arnCmd represents the arn command
var arnCmd = &cobra.Command{
	Use:   "arn [SUBSTR]",
	Short: "Get ARNs of load-balancers",
	Long:  `Get ARNs of load-balancers, optionally matching sub-string.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			getArns("")
		case 1:
			getArns(args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func init() {
	rootCmd.AddCommand(arnCmd)
}

func getArns(substr string) {
	for _, l := range loadBalancers(substr) {
		fmt.Println(*l.LoadBalancerArn)
	}
}
