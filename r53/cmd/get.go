package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get CMD",
	Aliases: []string{"g"},
	Short:   "Get route53 resources",
	Long:    `Get detail for route53 resources: hosted zones.`,
}

func init() {
	rootCmd.AddCommand(getCmd)
}
