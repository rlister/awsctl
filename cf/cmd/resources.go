package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/spf13/cobra"
	"log"
	"os"
	"text/tabwriter"
)

// resourcesCmd represents the resources command
var resourcesCmd = &cobra.Command{
	Use:     "resources STACK",
	Aliases: []string{"r"},
	Short:   "list stack resources",
	Long:    `List resources for stack with given name.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("wrong number of arguments")
		}
		listResources(args[0])
	},
}

// listResources lists resources for given stack
func listResources(name string) {
	res, err := client.DescribeStackResources(context.TODO(), &cloudformation.DescribeStackResourcesInput{
		StackName: &name,
	})

	if err != nil {
		log.Fatal(err)
	}

	const format = "%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "LOGICAL ID", "TYPE", "STATUS", "PHYSICAL ID")

	for _, r := range res.StackResources {
		fmt.Fprintf(tw, format, *r.LogicalResourceId, *r.ResourceType, statusColor(string(r.ResourceStatus)), *r.PhysicalResourceId)
	}

	tw.Flush()
}

func init() {
	rootCmd.AddCommand(resourcesCmd)
}
