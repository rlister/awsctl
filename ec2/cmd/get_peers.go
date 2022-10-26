package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

// peersCmd represents the peers command
var getPeersCmd = &cobra.Command{
	Use:     "peers",
	Aliases: []string{"peer", "p"},
	Short:   "Get ec2 peers",
	Long:    `Get ec2 peer details.`,
	Run: func(cmd *cobra.Command, args []string) {
		getPeers()
	},
}

func init() {
	getCmd.AddCommand(getPeersCmd)
}

func getPeers() {
	const format = "%v\t%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "NAME", "ID", "REQUESTER", "ACCEPTER", "STATUS")

	for _, p := range peers("") {
		name := findNameTag(p.Tags)
		fmt.Fprintf(tw, format, name, *p.VpcPeeringConnectionId, *p.RequesterVpcInfo.VpcId, *p.AccepterVpcInfo.VpcId, p.Status.Code)
	}

	tw.Flush()
}
