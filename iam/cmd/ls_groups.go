package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// groupsCmd represents the groups command
var groupsCmd = &cobra.Command{
	Use:     "groups [SUBSTR]",
	Aliases: []string{"g"},
	Short:   "List groups",
	Long:    `List iam groups, optionally matching string.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			listGroups("")
		case 1:
			listGroups(args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func init() {
	lsCmd.AddCommand(groupsCmd)
}

func listGroups(substr string) {
	for _, g := range groups(substr) {
		fmt.Println(*g.GroupName)
	}
}
