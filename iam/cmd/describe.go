package cmd

import (
	"github.com/spf13/cobra"
)

// describeCmd represents the describe command
var describeCmd = &cobra.Command{
	Use:     "describe",
	Aliases: []string{"d"},
	Short:   "Describe role or policy",
	Long:    `Describe given role or policy.`,
}

func init() {
	rootCmd.AddCommand(describeCmd)
}
