package cmd

import (
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
	"log"
	"context"
	"fmt"
	"time"
	"strconv"
	"strings"
)

const dateFormat = "2006-01-02 15:04 MST"

var getCmd = &cobra.Command{
	Use:     "get PREFIX",
	Aliases: []string{"g"},
	Short:   "get attributes for queues",
	Long:    `Get attributes of queues with prefix.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			getQueues("")
		default:
			getQueues(args[0])
		}
	},
}

func queueName(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}

func formatEpoch(timestamp string) string {
	epoch, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return time.Unix(epoch, 0).Format(dateFormat)
}

func getQueues(prefix string) {
	const format = "%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "NAME", "CREATED", "AVAIL", "IN-FLIGHT")

	attr := []types.QueueAttributeName{
		"CreatedTimestamp",
		"ApproximateNumberOfMessages",
		"ApproximateNumberOfMessagesNotVisible",
	}

	for _, q := range queues(prefix) {
		resp, err := client.GetQueueAttributes(context.TODO(), &sqs.GetQueueAttributesInput{
			AttributeNames: attr,
			QueueUrl: &q,
		})

		if err != nil {
			log.Fatal(err)
		}

		created := formatEpoch(resp.Attributes["CreatedTimestamp"])
		visible := resp.Attributes["ApproximateNumberOfMessages"]
		inflight := resp.Attributes["ApproximateNumberOfMessagesNotVisible"]
		fmt.Fprintf(tw, format, queueName(q), created, visible, inflight)
	}

	tw.Flush()
}

func init() {
	rootCmd.AddCommand(getCmd)
}
