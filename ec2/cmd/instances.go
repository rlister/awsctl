package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"log"
)

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

	// sort.Sort(byInstanceName(stacks))
	return instances
}
