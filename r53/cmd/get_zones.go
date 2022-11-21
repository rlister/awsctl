package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

// getZonesCmd represents the zones command
var getZonesCmd = &cobra.Command{
	Use:     "zones [SUBSTR]",
	Aliases: []string{"z"},
	Short:   "Get route53 zones",
	Long:    `Get route53 zones, optionally matching given sub-string.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			getZones("")
		default:
			getZones(args[0])
		}
	},
}

func init() {
	getCmd.AddCommand(getZonesCmd)
}

// getZones gets hosted zone details
func getZones(substr string) {
	const format = "%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "NAME", "ID", "COUNT", "DESC")
	for _, z := range zones(substr) {
		fmt.Fprintf(tw, format, *z.Name, *z.Id, *z.ResourceRecordSetCount, *z.Config.Comment)
	}
	tw.Flush()
}
