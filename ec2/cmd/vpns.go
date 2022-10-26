package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"log"
)

// return sorted slice of Vpns for vpns matching prefix
func vpns(prefix string) []types.ClientVpnEndpoint {
	paginator := ec2.NewDescribeClientVpnEndpointsPaginator(client, &ec2.DescribeClientVpnEndpointsInput{})

	vpns := []types.ClientVpnEndpoint{}

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range output.ClientVpnEndpoints {
			vpns = append(vpns, v)
		}
	}

	return vpns
}
