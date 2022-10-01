package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"log"
)

// return sorted slice of SecurityGroups for securitygroups matching prefix
func securityGroups(prefix string) []types.SecurityGroup {
	paginator := ec2.NewDescribeSecurityGroupsPaginator(client, &ec2.DescribeSecurityGroupsInput{})

	securityGroups := []types.SecurityGroup{}

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		for _, s := range output.SecurityGroups {
			securityGroups = append(securityGroups, s)
		}
	}

	// sort.Sort(bySecurityGroupName(stacks))
	return securityGroups
}
