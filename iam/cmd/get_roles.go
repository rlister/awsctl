package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"text/tabwriter"
	"os"
)

// getRolesCmd represents the roles command
var getRolesCmd = &cobra.Command{
	Use:   "roles [SUBSTR]",
	Aliases: []string{"r"},
	Short: "Get iam roles",
	Long:  `Get iam roles, optionally matching given sub-string.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			getRoles("")
		default:
			getRoles(args[0])
		}
	},
}

func init() {
	getCmd.AddCommand(getRolesCmd)
}

// getRoles gets role details
func getRoles(substr string) {
	const format = "%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "NAME", "CREATED", "ID")
	for _, r := range roles(substr) {
		fmt.Fprintf(tw, format, *r.RoleName, (*r.CreateDate).Format(dateFormat), *r.RoleId)
	}
	tw.Flush()
}
