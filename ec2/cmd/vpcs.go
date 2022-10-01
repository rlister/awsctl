package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"log"
)

// return sorted slice of Vpcs for vpcs matching prefix
func vpcs(prefix string) []types.Vpc {
	paginator := ec2.NewDescribeVpcsPaginator(client, &ec2.DescribeVpcsInput{})

	vpcs := []types.Vpc{}

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		for _, s := range output.Vpcs {
			vpcs = append(vpcs, s)
		}
	}

	// sort.Sort(byVpcName(stacks))
	return vpcs
}
