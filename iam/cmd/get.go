package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get CMD",
	Short: "Get iam resources",
	Long: `Get detail for iam resources: roles, profiles.`,
}

func init() {
	rootCmd.AddCommand(getCmd)
}
