package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// policiesCmd represents the policies command
var lsPoliciesCmd = &cobra.Command{
	Use:     "policies [SUBSTR]",
	Aliases: []string{"p"},
	Short:   "List iam policies",
	Long:    `List iam policies, optionally matching given sub-string.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			listPolicies("")
		default:
			listPolicies(args[0])
		}
	},
}

func init() {
	lsCmd.AddCommand(lsPoliciesCmd)
}

func listPolicies(substr string) {
	for _, p := range policies(substr) {
		fmt.Println(*p.PolicyName)
	}
}
