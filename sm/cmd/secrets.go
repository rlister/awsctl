package cmd

import (
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"log"
	"context"
	"sort"
	"strings"
)

// sort arrays of SecretListEntry
type byName []types.SecretListEntry

func (x byName) Len() int           { return len(x) }
func (x byName) Less(i, j int) bool { return strings.ToLower(*x[i].Name) < strings.ToLower(*x[j].Name) }
func (x byName) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func secrets(prefix string) []types.SecretListEntry {
	filters := []types.Filter{}
	if prefix != "" {
		filters = append(filters, types.Filter{Key: "name", Values: []string{prefix}})
	}

	paginator := secretsmanager.NewListSecretsPaginator(client, &secretsmanager.ListSecretsInput{Filters: filters})

	secrets := []types.SecretListEntry{}

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		for _, s := range output.SecretList {
			secrets = append(secrets, s)
		}
	}

	sort.Sort(byName(secrets))
	return secrets
}
