package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/spf13/cobra"
)

const dateFormat = "2006-01-02 15:04 MST"

// tablesCmd represents the tables command
var getTablesCmd = &cobra.Command{
	Use:     "tables [SUBSTR]",
	Aliases: []string{"t"},
	Short:   "Get tables details",
	Long:    `Get details for tables, optionally matching sub-string.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			s := ""
			getTables(&s)
		case 1:
			getTables(&args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func getTables(substr *string) {
	const format = "%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "NAME", "CREATED", "ITEMS")

	for _, name := range tables(substr) {
		output, err := client.DescribeTable(context.TODO(), &dynamodb.DescribeTableInput{TableName: &name})
		if err != nil {
			log.Fatal(err)
		}
		t := output.Table
		fmt.Fprintf(tw, format, *t.TableName, (*t.CreationDateTime).Format(dateFormat), t.ItemCount)
	}
	tw.Flush()
}

func init() {
	getCmd.AddCommand(getTablesCmd)
}
