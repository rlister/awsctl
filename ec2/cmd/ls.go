package cmd

import (
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List ec2 resources",
	Long: `List ec2 resources.`,
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
