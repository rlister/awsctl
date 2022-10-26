package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"log"
)

// return sorted slice of VpcPeeringConnections
func peers(prefix string) []types.VpcPeeringConnection {
	paginator := ec2.NewDescribeVpcPeeringConnectionsPaginator(client, &ec2.DescribeVpcPeeringConnectionsInput{})

	peers := []types.VpcPeeringConnection{}

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		for _, p := range output.VpcPeeringConnections {
			peers = append(peers, p)
		}
	}

	return peers
}
