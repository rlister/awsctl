package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get CMD",
	Aliases: []string{"g"},
	Short:   "Get dynamodb resources",
	Long:    `Get DynamoDB resources: tables.`,
}

func init() {
	rootCmd.AddCommand(getCmd)
}
