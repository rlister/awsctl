package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/spf13/cobra"
	"log"
)

// valueCmd represents the value command
var valueCmd = &cobra.Command{
	Use:     "value NAME",
	Aliases: []string{"v"},
	Short:   "Get secret value",
	Long:    `Get secret value for given secret name.`,
	Run: func(cmd *cobra.Command, args []string) {
		getValue(args[0])
	},
}

func getValue(name string) {
	output, err := client.GetSecretValue(context.TODO(), &secretsmanager.GetSecretValueInput{
		SecretId: &name,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(formatJson(output.SecretString))
}

// pretty-print json string
func formatJson(data *string) string {
	var buf bytes.Buffer
	err := json.Indent(&buf, []byte(*data), "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

func init() {
	rootCmd.AddCommand(valueCmd)
}
