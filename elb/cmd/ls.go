package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/spf13/cobra"
	"log"
	"sort"
	"strings"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls [SUBSTR]",
	Short: "List ELBs",
	Long:  `List all ELBs, optionally matching given sub-string.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			listElbs("")
		case 1:
			listElbs(args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}

// listLoadBalancers
func listElbs(substr string) {
	for _, l := range elbs(substr) {
		fmt.Println(*l.LoadBalancerName)
	}
}

// sort arrays of LoadBalancers
type byName []types.LoadBalancerDescription

func (x byName) Len() int           { return len(x) }
func (x byName) Less(i, j int) bool { return *x[i].LoadBalancerName < *x[j].LoadBalancerName }
func (x byName) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func elbs(substr string) []types.LoadBalancerDescription {
	paginator := elasticloadbalancing.NewDescribeLoadBalancersPaginator(client, &elasticloadbalancing.DescribeLoadBalancersInput{})

	lbs := []types.LoadBalancerDescription{}

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		for _, l := range output.LoadBalancerDescriptions {
			if strings.Contains(*l.LoadBalancerName, substr) {
				lbs = append(lbs, l)
			}
		}
	}

	sort.Sort(byName(lbs))
	return lbs
}
