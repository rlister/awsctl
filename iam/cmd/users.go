package cmd

import (

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"context"
	"log"
	"strings"
)

func users(substr string) []types.User {
	paginator := iam.NewListUsersPaginator(client, &iam.ListUsersInput{}, func(o *iam.ListUsersPaginatorOptions) {})

	users := []types.User{}

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		for _, r := range output.Users {
			if strings.Contains(*r.UserName, substr) {
				users = append(users, r)
			}
		}
	}
	return users
}
