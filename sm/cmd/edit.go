package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/spf13/cobra"
	"log"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:     "edit NAME",
	Aliases: []string{"e"},
	Short:   "Edit secret",
	Long:    `Edit secret with given name.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 1:
			editSecret(args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func editSecret(name string) {
	old, err := client.GetSecretValue(context.TODO(), &secretsmanager.GetSecretValueInput{SecretId: &name})
	if err != nil {
		log.Fatal(err)
	}

	data := edit(*old.SecretString)
	output, err := client.PutSecretValue(context.TODO(), &secretsmanager.PutSecretValueInput{SecretId: &name, SecretString: &data})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(*output.VersionId)
}

func init() {
	rootCmd.AddCommand(editCmd)
}
