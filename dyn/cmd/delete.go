package cmd

import (
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete CMD",
	Aliases: []string{"del"},
	Short:   "Delete dynamodb resources",
	Long:    `Delete dynamodb resources.`,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
