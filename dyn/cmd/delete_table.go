package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/spf13/cobra"
	"log"
)

// tablesCmd represents the tables command
var deleteTablesCmd = &cobra.Command{
	Use:     "table NAME",
	Aliases: []string{"t"},
	Short:   "Delete table",
	Long:    `Delete dynamodb table.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 1:
			deleteTable(&args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func deleteTable(name *string) {
	output, err := client.DeleteTable(context.TODO(), &dynamodb.DeleteTableInput{TableName: name})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*output.TableDescription.TableArn)
}

func init() {
	deleteCmd.AddCommand(deleteTablesCmd)
}
