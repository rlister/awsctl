package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// usersCmd represents the users command
var usersCmd = &cobra.Command{
	Use:   "users",
	Aliases: []string{"u"},
	Short: "List users",
	Long: `List iam users.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			listUsers("")
		case 1:
			listUsers(args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func init() {
	lsCmd.AddCommand(usersCmd)
}

func listUsers(substr string) {
	for _, r := range users(substr) {
		fmt.Println(*r.UserName)
	}
}
