package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls PREFIX",
	Short: "list queues",
	Long:  `List SQS queues matching PREFIX, or all if none given.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			listQueues("")
		default:
			listQueues(args[0])
		}
	},
}

func listQueues(prefix string) {
	for _, q := range queues(prefix) {
		fmt.Println(q)
	}
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
