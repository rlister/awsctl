package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/spf13/cobra"
	"log"
)

// versionsCmd represents the versions command
var versionsCmd = &cobra.Command{
	Use:     "versions [NAME]",
	Aliases: []string{"ver"},
	Short:   "List secrets versions",
	Long:    `List versions for given secret.`,
	Run: func(cmd *cobra.Command, args []string) {
		listVersions(args[0])
	},
}

func init() {
	rootCmd.AddCommand(versionsCmd)
}

func listVersions(name string) {
	output, err := client.ListSecretVersionIds(context.TODO(), &secretsmanager.ListSecretVersionIdsInput{SecretId: &name})
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range output.Versions {
		fmt.Println(v.CreatedDate, *v.VersionId, v.VersionStages)
	}
}
