package cmd

import (
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls CMD",
	Short: "List iam resources",
	Long: `List given iam resources: roles, profile, users.`,
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
