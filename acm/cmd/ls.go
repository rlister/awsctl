package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls [SUBSTR]",
	Short: "List certificates",
	Long:  `List certificates, optionally with domain name matching given substring.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			listCerts("")
		case 1:
			listCerts(args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}

func listCerts(substr string) {
	for _, c := range certs(substr) {
		fmt.Println(*c.CertificateArn)
	}
}

func certs(substr string) []types.CertificateSummary {
	paginator := acm.NewListCertificatesPaginator(client, &acm.ListCertificatesInput{}, func(o *acm.ListCertificatesPaginatorOptions) {})

	certs := []types.CertificateSummary{}

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		for _, c := range output.CertificateSummaryList {
			if strings.Contains(*c.DomainName, substr) || strings.Contains(*c.CertificateArn, substr) {
				certs = append(certs, c)
			}
		}
	}
	return certs
}
