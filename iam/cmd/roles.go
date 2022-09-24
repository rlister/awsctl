package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"log"
	"strings"
)

func roles(substr string) []types.Role {
	paginator := iam.NewListRolesPaginator(client, &iam.ListRolesInput{}, func(o *iam.ListRolesPaginatorOptions) {})

	roles := []types.Role{}

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		for _, r := range output.Roles {
			if strings.Contains(*r.RoleName, substr) {
				roles = append(roles, r)
			}
		}
	}
	return roles
}
