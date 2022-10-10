package cmd

import (
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls CMD",
	Short: "List dynamodb resources.",
	Long:  `List dynamodb resources: tables.`,
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
