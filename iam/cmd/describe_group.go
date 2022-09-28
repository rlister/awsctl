package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/spf13/cobra"
	"log"
	"sort"
)

// groupCmd represents the group command
var groupCmd = &cobra.Command{
	Use:     "group NAME",
	Aliases: []string{"g"},
	Short:   "Describe group",
	Long:    `Describe given group.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 1:
			describeGroup(args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func init() {
	describeCmd.AddCommand(groupCmd)
}

// describe group and policies it uses
func describeGroup(name string) {
	describeGroupDetails(&name)
}

// describe details for iam group
func describeGroupDetails(name *string) {
	output, err := client.GetGroup(context.TODO(), &iam.GetGroupInput{GroupName: name})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Arn: %s\n", *output.Group.Arn)
	fmt.Printf("CreateDate: %s\n", *output.Group.CreateDate)
	fmt.Printf("GroupId: %s\n", *output.Group.GroupId)
	fmt.Printf("Users:\n")

	// sort and print users
	names := []string{}
	for _, u := range output.Users {
		names = append(names, *u.UserName)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("  %s\n", name)
	}
}
