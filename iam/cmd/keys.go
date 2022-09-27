package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"log"
)

// keys takes ptr to user name and returns slice of AccessKeyMetadata
func keys(name *string) []types.AccessKeyMetadata {
	output, err := client.ListAccessKeys(context.TODO(), &iam.ListAccessKeysInput{UserName: name})
	if err != nil {
		log.Fatal(err)
	}

	return output.AccessKeyMetadata
}
