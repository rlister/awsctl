package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
	"text/tabwriter"
)

const dateFormat = "2006-01-02 15:04 MST"

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get [SUBSTR]",
	Aliases: []string{"g"},
	Short:   "Get keys",
	Long:    `Get keys, optionally matching substring.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			getKmsKeys("")
		case 1:
			getKmsKeys(args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func aliasMap() map[string]string {
	paginator := kms.NewListAliasesPaginator(client, &kms.ListAliasesInput{})

	var aliases = make(map[string]string)

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		for _, a := range output.Aliases {
			if !strings.HasPrefix(*a.AliasName, "alias/aws/") {
				aliases[*a.TargetKeyId] = *a.AliasName
			}
		}
	}

	return aliases
}

func getKmsKeys(substr string) {
	aliases := aliasMap()
	const format = "%v\t%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "ALIAS", "NAME", "ENABLED", "CREATED", "STATE")

	for _, k := range keys(substr) {
		output, err := client.DescribeKey(context.TODO(), &kms.DescribeKeyInput{KeyId: k.KeyId})
		if err != nil {
			log.Fatal(err)
		}

		d := output.KeyMetadata

		alias, exists := aliases[*d.KeyId]
		if !exists {
			alias = "-"
		}

		fmt.Fprintf(tw, format, alias, *d.KeyId, d.Enabled, (*d.CreationDate).Format(dateFormat), d.KeyState)
	}
	tw.Flush()
}

func init() {
	rootCmd.AddCommand(getCmd)
}
