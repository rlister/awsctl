package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"log"
)

// return sorted slice of Subnets for subnets matching prefix
func subnets(prefix string) []types.Subnet {
	paginator := ec2.NewDescribeSubnetsPaginator(client, &ec2.DescribeSubnetsInput{})

	subnets := []types.Subnet{}

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		for _, s := range output.Subnets {
			subnets = append(subnets, s)
		}
	}

	// sort.Sort(bySubnetName(stacks))
	return subnets
}
