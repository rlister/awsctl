package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// instancesCmd represents the instances command
var instancesCmd = &cobra.Command{
	Use:     "instances",
	Aliases: []string{"i"},
	Short:   "List ec2 instances",
	Long:    `List ec2 instance ids.`,
	Run: func(cmd *cobra.Command, args []string) {
		listInstances()
	},
}

func init() {
	lsCmd.AddCommand(instancesCmd)
}

func listInstances() {
	for _, i := range instances("") {
		fmt.Println(*i.InstanceId)
	}
}
