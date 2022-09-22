package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/spf13/cobra"
	"log"
)

// describeCmd represents the describe command
var describeCmd = &cobra.Command{
	Use:     "describe",
	Aliases: []string{"d"},
	Short:   "Describe load-balancer",
	Long:    `Show all load-balancer properties as json.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("wrong number of arguments")
		}
		describeLoadBalancer(args[0])
	},
}

func init() {
	rootCmd.AddCommand(describeCmd)
}

// describeLoad-Balancers prints all load-balancer info as json
func describeLoadBalancer(name string) {
	res, err := client.DescribeLoadBalancers(context.TODO(), &elasticloadbalancingv2.DescribeLoadBalancersInput{
		Names: []string{name},
	})

	if err != nil {
		log.Fatal(err)
	}

	for _, l := range res.LoadBalancers {
		data, err := json.MarshalIndent(l, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(data))
	}
}
