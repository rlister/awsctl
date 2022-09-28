package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

// policyCmd represents the policy command
var policyCmd = &cobra.Command{
	Use:     "policy ARN",
	Aliases: []string{"p"},
	Short:   "Describe policy",
	Long:    `Describe policy with given ARN. If given a name instead, attempt to find the first policy with that name (slower).`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 1:
			describePolicy(&args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func init() {
	describeCmd.AddCommand(policyCmd)
}

// describe policy and policy document
func describePolicy(arn *string) {
	if !strings.HasPrefix(*arn, "arn:") {
		arn = findPolicyByName(arn)
	}
	p := getPolicy(arn)
	describePolicyDetails(&p)
	describePolicyVersion(&p)
}

// returns iam.Policy
func getPolicy(name *string) types.Policy {
	output, err := client.GetPolicy(context.TODO(), &iam.GetPolicyInput{PolicyArn: name})
	if err != nil {
		log.Fatal(err)
	}
	return *output.Policy
}

// describe policy details
func describePolicyDetails(p *types.Policy) {
	fmt.Printf("Arn: %s\n", *p.Arn)
	fmt.Printf("Path: %s\n", *p.Path)
	fmt.Printf("CreateDate: %s\n", *p.CreateDate)
	fmt.Printf("UpdateDate: %s\n", *p.UpdateDate)
	fmt.Printf("AttachmentCount: %d\n", *p.AttachmentCount)
	fmt.Printf("DefaultVersionId: %s\n", *p.DefaultVersionId)
	fmt.Printf("Tags:\n")
	for _, tag := range p.Tags {
		fmt.Printf("  %s: %s\n", *tag.Key, *tag.Value)
	}
}

// print formatted json for policy document of default policy version
func describePolicyVersion(p *types.Policy) {
	fmt.Println("PolicyDocument:")

	output, err := client.GetPolicyVersion(context.TODO(), &iam.GetPolicyVersionInput{PolicyArn: p.Arn, VersionId: p.DefaultVersionId})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("   ", formatJson(output.PolicyVersion.Document))
}
