package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"log"
	"sort"
	"strings"
)

const dateFormat = "2006-01-02 15:04 MST"

type byPolicyName []types.Policy

func (x byPolicyName) Len() int           { return len(x) }
func (x byPolicyName) Less(i, j int) bool { return *x[i].PolicyName < *x[j].PolicyName }
func (x byPolicyName) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func policies(substr string) []types.Policy {
	paginator := iam.NewListPoliciesPaginator(client, &iam.ListPoliciesInput{}, func(o *iam.ListPoliciesPaginatorOptions) {})

	policies := []types.Policy{}

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		for _, p := range output.Policies {
			if strings.Contains(*p.PolicyName, substr) {
				policies = append(policies, p)
			}
		}
	}

	sort.Sort(byPolicyName(policies))
	return policies
}
