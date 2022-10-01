package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var cfg aws.Config
var client *ec2.Client

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ec2",
	Short: "ec2 commands",
	Long:  `ec2 commands for eg instances, subnets, vpcs.`,
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

	client = ec2.NewFromConfig(cfg)
}
