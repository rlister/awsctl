package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"context"
	"log"
	"strings"
	"errors"
)

// lsZonesCmd represents the zones command
var lsZonesCmd = &cobra.Command{
	Use:   "zones [SUBSTR]",
	Aliases: []string{"z"},
	Short: "List route53 hosted zones",
	Long:  `List route53 hosted zones, optionally matching given sub-string.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			listZones("")
		default:
			listZones(args[0])
		}
	},
}

func init() {
	lsCmd.AddCommand(lsZonesCmd)
}

// listZones lists host zone names
func listZones(substr string) {

	for _, r := range zones(substr) {
		fmt.Println(*r.Name)
	}
}

func zones(substr string) []types.HostedZone {
	paginator := route53.NewListHostedZonesPaginator(client, &route53.ListHostedZonesInput{}, func(o *route53.ListHostedZonesPaginatorOptions) {})

	zones := []types.HostedZone{}

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		for _, r := range output.HostedZones {
			if strings.Contains(*r.Name, substr) {
				zones = append(zones, r)
			}
		}
	}
	return zones
}

// return hosted zone id for zone with given name
func findZoneId(name string) (string, error) {
	// not a domain so probably a zone, just return it
	if !strings.Contains(name, ".") {
		return name, nil
	}

	// aws wants trailing period on name
	if !strings.HasSuffix(name, ".") {
		name = name + string(".")
	}

	output, err := client.ListHostedZonesByName(context.TODO(), &route53.ListHostedZonesByNameInput{DNSName: &name})
	if err != nil {
		return "", err
	}

	for _, z := range output.HostedZones {
		if *z.Name == name {
			return *z.Id, nil
		}
	}
	return "", errors.New("zone not found")
}
