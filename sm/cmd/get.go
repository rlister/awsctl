package cmd

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/spf13/cobra"
	"log"
	"os"
	"text/tabwriter"
)

const dateFormat = "2006-01-02 15:04 MST"

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get [PREFIX]",
	Aliases: []string{"g"},
	Short:   "Get secrets",
	Long:    `Get secrets, optionally matching prefix.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			getSecrets("")
		case 1:
			getSecrets(args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func getSecrets(prefix string) {
	const format = "%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "NAME", "CREATED", "ACCESSED", "PRIMARY")
	for _, s := range secrets(prefix) {
		fmt.Fprintf(tw, format, *s.Name, (*s.CreatedDate).Format(dateFormat), lastAccessedDate(s), primaryRegion(s))
	}
	tw.Flush()
}

func lastAccessedDate(s types.SecretListEntry) string {
	if s.LastAccessedDate == nil {
		return "-"
	} else {
		return (*s.LastAccessedDate).Format(dateFormat)
	}
}

func primaryRegion(s types.SecretListEntry) string {
	if s.PrimaryRegion == nil {
		return "-"
	} else {
		return *s.PrimaryRegion
	}
}

func init() {
	rootCmd.AddCommand(getCmd)
}
