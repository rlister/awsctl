package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/spf13/cobra"
	"log"
)

// describeCmd represents the describe command
var describeCmd = &cobra.Command{
	Use:     "describe",
	Aliases: []string{"d"},
	Short:   "describe stack",
	Long:    `Show all stack properties as json.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("wrong number of arguments")
		}
		describeStack(args[0])
	},
}

// describeStacks prints all stack info as json
func describeStack(name string) {
	res, err := client.DescribeStacks(context.TODO(), &cloudformation.DescribeStacksInput{
		StackName: &name,
	})

	if err != nil {
		log.Fatal(err)
	}

	for _, stack := range res.Stacks {
		data, err := json.MarshalIndent(stack, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(data))
	}
}

func init() {
	rootCmd.AddCommand(describeCmd)
}
