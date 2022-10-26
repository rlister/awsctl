package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/spf13/cobra"
	"log"
	"os"
	"text/tabwriter"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:     "user NAME",
	Aliases: []string{"u"},
	Short:   "Describe user",
	Long:    `Describe given user.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 1:
			describeUser(args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func init() {
	describeCmd.AddCommand(userCmd)
}

// describe user and policies it uses
func describeUser(name string) {
	describeUserDetails(&name)
	describeUserPolicies(&name)
	describeUserAttachedPolicies(&name)
	describeUserGroups(&name)
	describeUserKeys(&name)
}

// describe details for iam user
func describeUserDetails(name *string) {
	r, err := client.GetUser(context.TODO(), &iam.GetUserInput{UserName: name})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Arn: %s\n", *r.User.Arn)
	fmt.Printf("CreateDate: %s\n", *r.User.CreateDate)
	fmt.Printf("UserId: %s\n", *r.User.UserId)
	fmt.Printf("PasswordLastUsed: %s\n", lastUsed(*r.User))
}

// describe inline policies for user
func describeUserPolicies(name *string) {
	fmt.Println("Inline policies:")

	output, err := client.ListUserPolicies(context.TODO(), &iam.ListUserPoliciesInput{UserName: name})
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range output.PolicyNames {
		fmt.Println(p)
	}
}

// describe names of attached policies
func describeUserAttachedPolicies(name *string) {
	fmt.Println("Attached policies:")

	output, err := client.ListAttachedUserPolicies(context.TODO(), &iam.ListAttachedUserPoliciesInput{UserName: name})
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range output.AttachedPolicies {
		fmt.Printf("  %s\n", *p.PolicyName)
	}
}

func describeUserGroups(name *string) {
	fmt.Println("Groups:")

	output, err := client.ListGroupsForUser(context.TODO(), &iam.ListGroupsForUserInput{UserName: name})
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range output.Groups {
		fmt.Printf("  %s\n", *p.GroupName)
	}
}

func describeUserKeys(name *string) {
	fmt.Println("Access keys:")

	output, err := client.ListAccessKeys(context.TODO(), &iam.ListAccessKeysInput{UserName: name})
	if err != nil {
		log.Fatal(err)
	}

	const format = "  %v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "ID", "CREATED", "STATUS", "LAST USED")

	for _, k := range output.AccessKeyMetadata {
		l, err := client.GetAccessKeyLastUsed(context.TODO(), &iam.GetAccessKeyLastUsedInput{AccessKeyId: k.AccessKeyId})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(tw, format, *k.AccessKeyId, (*k.CreateDate).Format(dateFormat), k.Status, (*l.AccessKeyLastUsed.LastUsedDate).Format(dateFormat))
	}

	tw.Flush()
}
