package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"context"
	"log"
	"strings"
)

// lsRecordsCmd represents the zones command
var lsRecordsCmd = &cobra.Command{
	Use:   "records ZONE [SUBSTR]",
	Aliases: []string{"r"},
	Short: "List route53 records",
	Long:  `List route53 resource records sets for given zone.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 1:
			listRecords(args[0], "")
		case 2:
			listRecords(args[0], args[1])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func init() {
	lsCmd.AddCommand(lsRecordsCmd)
}

// listRecords lists host zone names
func listRecords(zone, substr string) {
	zone, err := findZoneId(zone)

	if err != nil {
		log.Fatal(err)
	}

	for _, r := range records(zone, substr, "") {
		fmt.Println(*r.Name)
	}
}

func records(zone, substr, typ string) []types.ResourceRecordSet {
	var max int32 = 300

	output, err := client.ListResourceRecordSets(context.TODO(), &route53.ListResourceRecordSetsInput{
		HostedZoneId: &zone,
		MaxItems: &max,
	})

	if err != nil {
		log.Fatal(err)
	}

	records := []types.ResourceRecordSet{}

	t := strings.ToUpper(typ)
	for _, r := range output.ResourceRecordSets {
		if strings.Contains(*r.Name, substr) && strings.Contains(string(r.Type), t) {
			records = append(records, r)
		}
	}

	return records
}
