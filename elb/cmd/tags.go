package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/spf13/cobra"
	"log"
	"text/tabwriter"
	"os"
)

// describeCmd represents the describe command
var tagsCmd = &cobra.Command{
	Use:     "tags",
	Aliases: []string{"t"},
	Short:   "Show ELB tags",
	Long:    `Show all ELB tags.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("wrong number of arguments")
		}
		describeTags(args[0])
	},
}

func init() {
	rootCmd.AddCommand(tagsCmd)
}

// describeLoad-Balancers prints all load-balancer info as json
func describeTags(name string) {
	res, err := client.DescribeTags(context.TODO(), &elasticloadbalancing.DescribeTagsInput{
		LoadBalancerNames: []string{name},
	})

	if err != nil {
		log.Fatal(err)
	}

	const format = "%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)

	for _, d := range res.TagDescriptions {
		for _, t := range d.Tags {
			fmt.Fprintf(tw, format, *t.Key, *t.Value)
		}
	}

	tw.Flush()
}
