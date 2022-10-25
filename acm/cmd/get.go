package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"
	"github.com/spf13/cobra"
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
)

const dateFormat = "2006-01-02 15:04 MST"

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get [SUBSTR]",
	Aliases: []string{"g"},
	Short:   "List certificate summaries",
	Long:    `List certificate summaries, optionally matching given domain substring.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			getCerts("")
		case 1:
			getCerts(args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func getCerts(substr string) {
	const format = "%v\t%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "DOMAIN", "ID", "IN USE", "NOT AFTER", "STATUS")

	for _, c := range certs(substr) {
		id := strings.SplitAfter(*c.CertificateArn, "/")[1]
		fmt.Fprintf(tw, format, *c.DomainName, id, *c.InUse, notAfter(c), statusColor(string(c.Status)))
	}
	tw.Flush()
}

func notAfter(c types.CertificateSummary) string {
	if c.NotAfter == nil {
		return "-"
	}
	return (*c.NotAfter).Format(dateFormat)
}
