package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// subnetsCmd represents the subnets command
var subnetsCmd = &cobra.Command{
	Use:     "subnets",
	Aliases: []string{"s"},
	Short:   "List ec2 subnets",
	Long:    `List ec2 subnet ids.`,
	Run: func(cmd *cobra.Command, args []string) {
		listSubnets()
	},
}

func init() {
	lsCmd.AddCommand(subnetsCmd)
}

func listSubnets() {
	for _, s := range subnets("") {
		fmt.Println(*s.SubnetId)
	}
}
