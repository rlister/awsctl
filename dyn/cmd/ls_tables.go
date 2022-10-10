package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

// tablesCmd represents the tables command
var tablesCmd = &cobra.Command{
	Use:     "tables",
	Aliases: []string{"t"},
	Short:   "List dynamodb tables",
	Long:    `List dynamodb tables.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			s := ""
			listTables(&s)
		case 1:
			listTables(&args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func listTables(substr *string) {
	for _, t := range tables(substr) {
		fmt.Println(t)
	}
}

func tables(substr *string) []string {
	paginator := dynamodb.NewListTablesPaginator(client, &dynamodb.ListTablesInput{})

	tables := []string{}

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		for _, n := range output.TableNames {
			if strings.Contains(n, *substr) {
				tables = append(tables, n)
			}
		}
	}
	return tables
}

func init() {
	lsCmd.AddCommand(tablesCmd)
}
