package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

// getUsersCmd represents the users command
var getUsersCmd = &cobra.Command{
	Use:     "users [SUBSTR]",
	Aliases: []string{"u"},
	Short:   "Get iam users",
	Long:    `Get iam users, optionally matching given sub-string.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			getUsers("")
		default:
			getUsers(args[0])
		}
	},
}

func init() {
	getCmd.AddCommand(getUsersCmd)
}

// getUsers gets user details
func getUsers(substr string) {
	const format = "%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "NAME",  "ID", "CREATED", "LAST USED")
	for _, u := range users(substr) {
		fmt.Fprintf(tw, format, *u.UserName, *u.UserId, (*u.CreateDate).Format(dateFormat), lastUsed(u))
	}
	tw.Flush()
}

func lastUsed(u types.User) string {
	if u.PasswordLastUsed == nil {
		return "-"
	} else {
		return (*u.PasswordLastUsed).Format(dateFormat)
	}
}
