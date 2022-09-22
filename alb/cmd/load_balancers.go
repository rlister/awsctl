package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"log"
	"sort"
	"strings"
)

// sort arrays of LoadBalancers
type byName []types.LoadBalancer

func (x byName) Len() int           { return len(x) }
func (x byName) Less(i, j int) bool { return *x[i].LoadBalancerName < *x[j].LoadBalancerName }
func (x byName) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func loadBalancers(substr string) []types.LoadBalancer {
	paginator := elasticloadbalancingv2.NewDescribeLoadBalancersPaginator(client, &elasticloadbalancingv2.DescribeLoadBalancersInput{})

	lbs := []types.LoadBalancer{}

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		for _, l := range output.LoadBalancers {
			if strings.Contains(*l.LoadBalancerName, substr) {
				lbs = append(lbs, l)
			}
		}
	}

	sort.Sort(byName(lbs))
	return lbs
}
