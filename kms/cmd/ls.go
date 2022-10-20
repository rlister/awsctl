package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List kms keys",
	Long:  `List kms keys.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			listKmsKeys("")
		case 1:
			listKmsKeys(args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func listKmsKeys(substr string) {
	for _, k := range keys(substr) {
		fmt.Println(*k.KeyId)
	}
}

func keys(substr string) []types.KeyListEntry {
	paginator := kms.NewListKeysPaginator(client, &kms.ListKeysInput{})

	keys := []types.KeyListEntry{}

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		for _, k := range output.Keys {
			if strings.Contains(*k.KeyId, substr) {
				keys = append(keys, k)
			}
		}
	}
	return keys
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
