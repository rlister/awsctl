package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var cfg aws.Config
var client *elasticloadbalancingv2.Client

var rootCmd = &cobra.Command{
	Use:   "alb",
	Short: "alb commands",
	Long:  `ALB commands.`,
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

	client = elasticloadbalancingv2.NewFromConfig(cfg)
}
