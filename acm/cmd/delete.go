package cmd

import (
	"bufio"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete ID",
	Aliases: []string{"del"},
	Short:   "Delete cert",
	Long:    `Delete certificate with given id.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("wrong number of arguments")
		}
		deleteCert(args[0])
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

// delete matching certs if not in use
func deleteCert(id string) {
	input := bufio.NewScanner(os.Stdin)

	for _, c := range certs(id) {
		if *c.InUse {
			fmt.Println(*c.CertificateArn, "in use: not deleting")
			break
		}

		fmt.Printf("Delete %s? [y/n] ", *c.CertificateArn)
		input.Scan()

		if input.Text() == "y" {
			_, err := client.DeleteCertificate(context.TODO(), &acm.DeleteCertificateInput{CertificateArn: c.CertificateArn})
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
