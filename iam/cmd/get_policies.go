package cmd

import (
	"github.com/spf13/cobra"
	"text/tabwriter"
	"fmt"
	"os"
)

var getPoliciesCmd = &cobra.Command{
	Use:     "policies [SUBSTR]",
	Aliases: []string{"p"},
	Short:   "Get iam policies",
	Long:    `Get iam policies, optionally matching given sub-string.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			getPolicies("")
		default:
			getPolicies(args[0])
		}
	},
}

func init() {
	getCmd.AddCommand(getPoliciesCmd)
}

func getPolicies(substr string) {
	const format = "%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "NAME", "CREATED", "UPDATED", "ATTACH")
	for _, p := range policies(substr) {
		fmt.Fprintf(tw, format, *p.PolicyName, (*p.CreateDate).Format(dateFormat), (*p.UpdateDate).Format(dateFormat), *p.AttachmentCount)
	}
	tw.Flush()
}
