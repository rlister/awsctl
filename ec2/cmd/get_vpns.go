package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

// vpnsCmd represents the vpns command
var getVpnsCmd = &cobra.Command{
	Use:     "vpns",
	Aliases: []string{"vpn"},
	Short:   "Get vpns endpoints",
	Long:    `Get vpn endpoint details.`,
	Run: func(cmd *cobra.Command, args []string) {
		getVpns()
	},
}

func init() {
	getCmd.AddCommand(getVpnsCmd)
}

func getVpns() {
	const format = "%v\t%v\t%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "NAME", "ID", "CREATED", "CIDR", "VPC", "STATUS")
	for _, v := range vpns("") {
		name := findNameTag(v.Tags)
		fmt.Fprintf(tw, format, name, *v.ClientVpnEndpointId, *v.CreationTime, *v.ClientCidrBlock, *v.VpcId, *&v.Status.Code)
	}
	tw.Flush()
}
