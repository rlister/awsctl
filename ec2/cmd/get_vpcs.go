package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

// vpcsCmd represents the vpcs command
var getVpcsCmd = &cobra.Command{
	Use:     "vpcs",
	Aliases: []string{"v"},
	Short:   "Get ec2 vpcs",
	Long:    `Get ec2 vpc details.`,
	Run: func(cmd *cobra.Command, args []string) {
		getVpcs()
	},
}

func init() {
	getCmd.AddCommand(getVpcsCmd)
}

func getVpcs() {
	const format = "%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "NAME", "ID", "CIDR")
	for _, v := range vpcs("") {
		name := findNameTag(v.Tags)
		fmt.Fprintf(tw, format, name, *v.VpcId, *v.CidrBlock)
	}
	tw.Flush()
}
