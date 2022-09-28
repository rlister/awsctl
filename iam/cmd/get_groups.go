package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

// getGroupsCmd represents the groups command
var getGroupsCmd = &cobra.Command{
	Use:     "groups [SUBSTR]",
	Aliases: []string{"g"},
	Short:   "Get iam groups",
	Long:    `Get iam groups, optionally matching given sub-string.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			getGroups("")
		default:
			getGroups(args[0])
		}
	},
}

func init() {
	getCmd.AddCommand(getGroupsCmd)
}

// getGroups gets user details
func getGroups(substr string) {
	const format = "%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "NAME", "PATH", "ID", "CREATED")
	for _, g := range groups(substr) {
		fmt.Fprintf(tw, format, *g.GroupName, *g.Path, *g.GroupId, (*g.CreateDate).Format(dateFormat))
	}
	tw.Flush()
}
