package cmd

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/spf13/cobra"
	"os"
	"sort"
	"text/tabwriter"
)

// securityGroupsCmd represents the securityGroups command
var getSecurityGroupsCmd = &cobra.Command{
	Use:     "sg",
	Aliases: []string{"s"},
	Short:   "Get ec2 securityGroups",
	Long:    `Get ec2 securityGroup details.`,
	Run: func(cmd *cobra.Command, args []string) {
		getSecurityGroups()
	},
}

func init() {
	getCmd.AddCommand(getSecurityGroupsCmd)
}

// sort arrays of SecurityGroup
type byGroupName []types.SecurityGroup

func (x byGroupName) Len() int           { return len(x) }
func (x byGroupName) Less(i, j int) bool { return *x[i].GroupName < *x[j].GroupName }
func (x byGroupName) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func getSecurityGroups() {
	groups := securityGroups("")
	sort.Sort(byGroupName(groups))

	const format = "%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "NAME", "ID", "VPC")

	for _, s := range groups {
		fmt.Fprintf(tw, format, *s.GroupName, *s.GroupId, *s.VpcId)
	}
	tw.Flush()
}
