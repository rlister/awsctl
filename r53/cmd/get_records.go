package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"text/tabwriter"
)

// getRecordsCmd represents the records command
var getRecordsCmd = &cobra.Command{
	Use:     "records ZONE [SUBSTR]",
	Aliases: []string{"r"},
	Short:   "Get route53 records",
	Long:    `Get route53 records for given zone.`,
	Run: func(cmd *cobra.Command, args []string) {
		t, _ := cmd.Flags().GetString("type")
		switch len(args) {
		case 1:
			getRecords(args[0], "", t)
		case 2:
			getRecords(args[0], args[1], t)
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func init() {
	getCmd.AddCommand(getRecordsCmd)
	getRecordsCmd.Flags().StringP("type", "t", "", "Limit record types")
}

// getRecords gets hosted record details
func getRecords(zone, substr, t string) {
	const format = "%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "NAME", "TYPE", "VALUE/ALIAS")
	for _, r := range records(zone, substr, t) {

		// handle alias records
		if r.AliasTarget != nil {
			fmt.Fprintf(tw, format, *r.Name, r.Type, *r.AliasTarget.DNSName)
		}

		// handle non-alias records
		for _, rr := range r.ResourceRecords {
			fmt.Fprintf(tw, format, *r.Name, r.Type, *rr.Value)
		}
	}
	tw.Flush()
}
