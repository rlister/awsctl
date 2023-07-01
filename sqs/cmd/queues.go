package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"log"
)

func queues(prefix string) []string {
	paginator := sqs.NewListQueuesPaginator(client, &sqs.ListQueuesInput{
		QueueNamePrefix: &prefix,
	})

	queues := []string{}

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		for _, q := range output.QueueUrls {
			queues = append(queues, q)
		}
	}

	return queues
}
