package cmd

import (
	"github.com/spf13/cobra"
)

// describeCmd represents the describe command
var describeCmd = &cobra.Command{
	Use:   "describe",
	Aliases: []string{"d"},
	Short: "Describe resource",
	Long: `Describe given resource.`,
}

func init() {
	rootCmd.AddCommand(describeCmd)
}
