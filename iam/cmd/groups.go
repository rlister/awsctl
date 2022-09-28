package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"log"
	"strings"
)

func groups(substr string) []types.Group {
	paginator := iam.NewListGroupsPaginator(client, &iam.ListGroupsInput{}, func(o *iam.ListGroupsPaginatorOptions) {})

	groups := []types.Group{}

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		for _, g := range output.Groups {
			if strings.Contains(*g.GroupName, substr) {
				groups = append(groups, g)
			}
		}
	}
	return groups
}
