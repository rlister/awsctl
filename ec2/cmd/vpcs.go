package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"log"
	"sort"
)

type byVpcName []types.Vpc

func (x byVpcName) Len() int           { return len(x) }
func (x byVpcName) Less(i, j int) bool { return findNameTag(x[i].Tags) < findNameTag(x[j].Tags) }
func (x byVpcName) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

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

	sort.Sort(byVpcName(vpcs))
	return vpcs
}
