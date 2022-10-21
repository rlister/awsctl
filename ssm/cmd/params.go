package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"log"
	"sort"
)

func init() {
	rootCmd.AddCommand(lsCmd)
}

// sort arrays of Parameter
type byName []types.Parameter

func (x byName) Len() int           { return len(x) }
func (x byName) Less(i, j int) bool { return *x[i].Name < *x[j].Name }
func (x byName) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// returns array of Parameters with given path
func paramsByPath(path string) []types.Parameter {
	recurse := true
	paginator := ssm.NewGetParametersByPathPaginator(client, &ssm.GetParametersByPathInput{Path: &path, Recursive: &recurse})

	params := []types.Parameter{}

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		for _, p := range output.Parameters {
			params = append(params, p)
		}
	}

	sort.Sort(byName(params))
	return params
}

// return single exact-match Parameter, or nil if not found
func paramByName(name string) *types.Parameter {
	output, err := client.GetParameter(context.TODO(), &ssm.GetParameterInput{Name: &name})
	if err != nil {
		return nil
	}
	return output.Parameter
}
