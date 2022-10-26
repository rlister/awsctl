package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// vpnsCmd represents the vpns command
var vpnsCmd = &cobra.Command{
	Use:     "vpns",
	Aliases: []string{"vpn"},
	Short:   "List vpn endpoints",
	Long:    `List vpn endpoints.`,
	Run: func(cmd *cobra.Command, args []string) {
		listVpns()
	},
}

func init() {
	lsCmd.AddCommand(vpnsCmd)
}

func listVpns() {
	for _, v := range vpns("") {
		fmt.Println(*v.ClientVpnEndpointId)
	}
}
