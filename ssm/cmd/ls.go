package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/spf13/cobra"
	"log"
	"time"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls [PREFIX]",
	Short: "List parameters",
	Long:  `List ssm parameters, optionally with prefix. This is best effort, and may be throttled.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			listParams("")
		case 1:
			listParams(args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func listParams(prefix string) {
	filters := []types.ParametersFilter{}
	if prefix != "" {
		filters = append(filters, types.ParametersFilter{Key: types.ParametersFilterKeyName, Values: []string{prefix}})
	}

	paginator := ssm.NewDescribeParametersPaginator(client, &ssm.DescribeParametersInput{Filters: filters})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		for _, p := range output.Parameters {
			fmt.Println(*p.Name)
		}

		// this api throttles hard
		time.Sleep(100 * time.Millisecond)
	}
}
