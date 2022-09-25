package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var cfg aws.Config
var client *ecr.Client

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ecr",
	Short: "ecr commands",
	Long:  `ECR commands.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
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

	client = ecr.NewFromConfig(cfg)
}
