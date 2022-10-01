package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// vpcsCmd represents the vpcs command
var vpcsCmd = &cobra.Command{
	Use:     "vpcs",
	Aliases: []string{"v"},
	Short:   "List ec2 vpcs",
	Long:    `List ec2 vpc ids.`,
	Run: func(cmd *cobra.Command, args []string) {
		listVpcs()
	},
}

func init() {
	lsCmd.AddCommand(vpcsCmd)
}

func listVpcs() {
	for _, s := range vpcs("") {
		fmt.Println(*s.VpcId)
	}
}
