package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"log"
	"sort"
)

// sort arrays of Repository
type byName []types.Repository

func (x byName) Len() int           { return len(x) }
func (x byName) Less(i, j int) bool { return *x[i].RepositoryName < *x[j].RepositoryName }
func (x byName) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func repos() []types.Repository {
	paginator := ecr.NewDescribeRepositoriesPaginator(client, &ecr.DescribeRepositoriesInput{})

	repos := []types.Repository{}

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		for _, r := range output.Repositories {
			repos = append(repos, r)
		}
	}
	sort.Sort(byName(repos))
	return repos
}
