package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

var describeCmd = &cobra.Command{
	Use:     "describe",
	Aliases: []string{"d"},
	Short:   "describe queue",
	Long:    `Describe all queue attributes as json.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("wrong number of arguments")
		}
		describeQueue(args[0])
	},
}

func describeQueue(name string) {
	resp, err := client.GetQueueAttributes(context.TODO(), &sqs.GetQueueAttributesInput{
		AttributeNames: []types.QueueAttributeName{"All"},
		QueueUrl: &name,
	})

	if err != nil {
		log.Fatal(err)
	}

	resp.Attributes["CreatedTimestamp"] = formatEpoch(resp.Attributes["CreatedTimestamp"])

	data, err := json.MarshalIndent(resp.Attributes, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}

func init() {
	rootCmd.AddCommand(describeCmd)
}
