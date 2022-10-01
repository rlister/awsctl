package cmd

import (
	"github.com/spf13/cobra"
)

const dateFormat = "2006-01-02 15:04 MST"

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Aliases: []string{"g"},
	Short: "Get ec2 resources",
	Long: `Get ec2 resources: instances, subnets, vpcs.`,
}

func init() {
	rootCmd.AddCommand(getCmd)
}
