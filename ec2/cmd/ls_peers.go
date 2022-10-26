package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// peersCmd represents the peers command
var peersCmd = &cobra.Command{
	Use:     "peers",
	Aliases: []string{"p"},
	Short:   "List ec2 peers",
	Long:    `List ec2 peer ids.`,
	Run: func(cmd *cobra.Command, args []string) {
		listPeers()
	},
}

func init() {
	lsCmd.AddCommand(peersCmd)
}

func listPeers() {
	for _, p := range peers("") {
		fmt.Println(*p.VpcPeeringConnectionId)
	}
}
