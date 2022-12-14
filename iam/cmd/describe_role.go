package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"context"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"net/url"
	"bytes"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

// roleCmd represents the role command
var roleCmd = &cobra.Command{
	Use:   "role",
	Aliases: []string{"r"},
	Short: "Describe role",
	Long: `Describe given role.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 1:
			describeRole(args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func init() {
	describeCmd.AddCommand(roleCmd)
}

// describe role and policies it uses
func describeRole(name string) {
	describeRoleDetails(&name)
	describeRolePolicies(&name)
	describeAttachedPolicies(&name)
}

// describe role details
func describeRoleDetails(name *string) {
	output, err := client.GetRole(context.TODO(), &iam.GetRoleInput{RoleName: name})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Arn: %s\n", *output.Role.Arn)
	fmt.Printf("CreateDate: %s\n", *output.Role.CreateDate)
	date, region := roleLastUsed(output.Role)
	fmt.Printf("RoleLastUsed: %s %s\n", date, region)
	fmt.Printf("AssumeRolePolicyDocument:\n    %s\n", formatJson(output.Role.AssumeRolePolicyDocument))
}

// takes a Role and returns last used date and region, or strings representing null values
func roleLastUsed(role *types.Role) (string, string) {
	date, region := "-", ""
	if role.RoleLastUsed.LastUsedDate != nil {
		date = (*role.RoleLastUsed.LastUsedDate).Format(dateFormat)
	}
	if role.RoleLastUsed.Region != nil {
		region = *role.RoleLastUsed.Region
	}
	return date, region
}

// describe inline policies for role
func describeRolePolicies(name *string) {
	fmt.Println("Inline policies:")

	r, err := client.ListRolePolicies(context.TODO(), &iam.ListRolePoliciesInput{RoleName: name})
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range r.PolicyNames {
		describeRolePolicy(name, &p)
	}
}

// describe an inline policy for role
func describeRolePolicy(role *string, policy *string) {
	fmt.Printf("  %s:\n", *policy)

	r, err := client.GetRolePolicy(context.TODO(), &iam.GetRolePolicyInput{RoleName: role, PolicyName: policy})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("   ", formatJson(r.PolicyDocument))
}

// describe names of attached policies
func describeAttachedPolicies(name *string) {
	fmt.Println("Attached policies:")

	r, err := client.ListAttachedRolePolicies(context.TODO(), &iam.ListAttachedRolePoliciesInput{RoleName: name})
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range r.AttachedPolicies {
		fmt.Printf("  %s\n", *p.PolicyArn)
	}
}

// url-escape and pretty-print embedded json string
func formatJson(str *string) string {
	doc, err := url.QueryUnescape(*str)
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer

	err = json.Indent(&buf, []byte(doc), "    ", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return buf.String()
}
