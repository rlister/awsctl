package cmd

import (
	"os"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/spf13/cobra"
	"log"
	"context"
)

var cfg aws.Config
var client *sqs.Client

var rootCmd = &cobra.Command{
	Use:   "sqs",
	Short: "sqs",
	Long: `SQS tools.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	var err error
	cfg, err = config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	client = sqs.NewFromConfig(cfg)
}
