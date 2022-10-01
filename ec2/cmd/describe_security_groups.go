package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/spf13/cobra"
	"log"
	"os"
	"text/tabwriter"
)

// securityGroupsCmd represents the securityGroups command
var describeSecurityGroupCmd = &cobra.Command{
	Use:     "sg",
	Aliases: []string{"s"},
	Short:   "Describe ec2 securityGroup",
	Long:    `Describe ec2 securityGroup details.`,
	Run: func(cmd *cobra.Command, args []string) {
		describeSecurityGroup(&args[0])
	},
}

func init() {
	describeCmd.AddCommand(describeSecurityGroupCmd)
}

func getGroup(id *string) *types.SecurityGroup {
	output, err := client.DescribeSecurityGroups(context.TODO(), &ec2.DescribeSecurityGroupsInput{GroupIds: []string{*id}})

	if err != nil {
		log.Fatal(err)
	}

	return &output.SecurityGroups[0]
}

func describeSecurityGroup(id *string) {
	group := getGroup(id)
	describeGroupDetails(group)
	describeIngress(group)
	describeEgress(group)
}

func describeGroupDetails(g *types.SecurityGroup) {
	fmt.Printf("GroupName: %s\n", *g.GroupName)
	fmt.Printf("OwnerId: %s\n", *g.OwnerId)
	fmt.Printf("VpcId: %s\n", *g.VpcId)
}

func describeIngress(g *types.SecurityGroup) {
	fmt.Println("\nInbound Rules:")
	describeIpPermissions(g.IpPermissions)
}

func describeEgress(g *types.SecurityGroup) {
	fmt.Println("\nOutbound Rules:")
	describeIpPermissions(g.IpPermissionsEgress)
}

func describeIpPermissions(ip []types.IpPermission) {
	const format = "  %v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "PROTOCOL", "PORTS", "SOURCE", "DESCRIPTON")

	for _, i := range ip {
		ports := portRange(&i)
		for _, r := range i.IpRanges {
			fmt.Fprintf(tw, format, *proto(i.IpProtocol), *ports, *r.CidrIp, *desc(r.Description))
		}
		for _, u := range i.UserIdGroupPairs {
			fmt.Fprintf(tw, format, *proto(i.IpProtocol), *ports, *u.GroupId, *desc(u.Description))
		}
	}
	tw.Flush()
}

// check string ptr and return empty string if nil
func desc(s *string) *string {
	str := ""
	if s != nil {
		str = *s
	}
	return &str
}

// human-readable protocol
func proto(s *string) *string {
	var str string
	switch *s {
	case "-1":
		str = "all"
	default:
		str = *s
	}
	return &str
}

// check from and to ports for nil, and print as readable range
func portRange(p *types.IpPermission) *string {
	ports, from, to := "", "all", "all"
	if p.FromPort != nil {
		from = fmt.Sprintf("%d", *p.FromPort)
	}
	if p.ToPort != nil {
		to = fmt.Sprintf("%d", *p.ToPort)
	}
	if from == to {
		ports = from
	} else {
		ports = fmt.Sprintf("%s-%s", from, to)
	}
	return &ports
}
