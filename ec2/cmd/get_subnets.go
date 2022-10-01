package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

// subnetsCmd represents the subnets command
var getSubnetsCmd = &cobra.Command{
	Use:     "subnets",
	Aliases: []string{"s"},
	Short:   "Get ec2 subnets",
	Long:    `Get ec2 subnet details.`,
	Run: func(cmd *cobra.Command, args []string) {
		getSubnets()
	},
}

func init() {
	getCmd.AddCommand(getSubnetsCmd)
}

func getSubnets() {
	const format = "%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "NAME", "ID", "VPC", "CIDR")
	for _, s := range subnets("") {
		name := findNameTag(s.Tags)
		fmt.Fprintf(tw, format, name, *s.SubnetId, *s.VpcId, *s.CidrBlock)
	}
	tw.Flush()
}
