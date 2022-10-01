package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/spf13/cobra"
)

// instancesCmd represents the instances command
var getInstancesCmd = &cobra.Command{
	Use:     "instances",
	Aliases: []string{"i"},
	Short:   "Get ec2 instances",
	Long:    `Get ec2 instance ids.`,
	Run: func(cmd *cobra.Command, args []string) {
		getInstances()
	},
}

func init() {
	getCmd.AddCommand(getInstancesCmd)
}

func getInstances() {
	const format = "%v\t%v\t%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "NAME", "ID", "TYPE", "AZ", "LAUNCH", "STATE")
	for _, i := range instances("") {
		name := findNameTag(i.Tags)
		fmt.Fprintf(tw, format, name, *i.InstanceId, *&i.InstanceType, *i.Placement.AvailabilityZone, (*i.LaunchTime).Format(dateFormat), statusColor(string(i.State.Name)))
	}
	tw.Flush()
}

// take array of Tags and return Name tag value if it exists
func findNameTag(tags []types.Tag) string {
	name := "-"
	for _, t := range tags {
		if *t.Key == "Name" {
			name = *t.Value
			break
		}
	}
	return name
}
