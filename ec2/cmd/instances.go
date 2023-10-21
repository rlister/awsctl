package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"log"
	"sort"
)

type byInstanceName []types.Instance

func (x byInstanceName) Len() int           { return len(x) }
func (x byInstanceName) Less(i, j int) bool { return findNameTag(x[i].Tags) < findNameTag(x[j].Tags) }
func (x byInstanceName) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// return sorted slice of Instances for instances matching prefix
func instances(prefix string) []types.Instance {
	paginator := ec2.NewDescribeInstancesPaginator(client, &ec2.DescribeInstancesInput{})

	instances := []types.Instance{}

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		for _, r := range output.Reservations {
			for _, i := range r.Instances {
				instances = append(instances, i)
			}
		}
	}

	sort.Sort(byInstanceName(instances))
	return instances
}
