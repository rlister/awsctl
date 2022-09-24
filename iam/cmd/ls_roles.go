package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// lsRolesCmd represents the roles command
var lsRolesCmd = &cobra.Command{
	Use:   "roles [SUBSTR]",
	Aliases: []string{"r"},
	Short: "List iam roles",
	Long:  `List iam roles, optionally matching given sub-string.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			listRoles("")
		default:
			listRoles(args[0])
		}
	},
}

func init() {
	lsCmd.AddCommand(lsRolesCmd)
}

// listRoles lists role names
func listRoles(substr string) {

	for _, r := range roles(substr) {
		fmt.Println(*r.RoleName)
	}
}
