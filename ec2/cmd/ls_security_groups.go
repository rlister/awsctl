package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// securityGroupsCmd represents the securityGroups command
var securityGroupsCmd = &cobra.Command{
	Use:     "sg",
	Aliases: []string{"s"},
	Short:   "List ec2 security groups",
	Long:    `List ec2 security group ids.`,
	Run: func(cmd *cobra.Command, args []string) {
		listSecurityGroups()
	},
}

func init() {
	lsCmd.AddCommand(securityGroupsCmd)
}

func listSecurityGroups() {
	for _, s := range securityGroups("") {
		fmt.Println(*s.GroupId)
	}
}
