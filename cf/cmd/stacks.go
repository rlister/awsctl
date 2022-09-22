package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"log"
	"sort"
	"strings"
)

// sort arrays of StackSummary
type byStackName []types.StackSummary

func (x byStackName) Len() int           { return len(x) }
func (x byStackName) Less(i, j int) bool { return *x[i].StackName < *x[j].StackName }
func (x byStackName) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// return sorted slice of StackSummaries for stacks matching prefix
func stacks(prefix string) []types.StackSummary {
	paginator := cloudformation.NewListStacksPaginator(client, &cloudformation.ListStacksInput{
		StackStatusFilter: existingStackStatuses(),
	})

	stacks := []types.StackSummary{}

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		for _, s := range output.StackSummaries {
			if strings.HasPrefix(*s.StackName, prefix) {
				stacks = append(stacks, s)
			}
		}
	}

	sort.Sort(byStackName(stacks))
	return stacks
}
