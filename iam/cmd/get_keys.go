package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"text/tabwriter"
	"os"
	"strings"
)

// keysCmd represents the users command
var keysCmd = &cobra.Command{
	Use:     "keys",
	Aliases: []string{"k"},
	Short:   "Get keys",
	Long:    `Get iam keys.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			getKeys("")
		case 1:
			getKeys(args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func init() {
	getCmd.AddCommand(keysCmd)
}

func getKeys(substr string) {
	const format = "%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "NAME", "KEYS")

	for _, u := range users(substr) {
		keyids := []string{}
		for _, k := range keys(u.UserName) {
			keyids = append(keyids, *k.AccessKeyId)
		}
		fmt.Fprintf(tw, format, *u.UserName, strings.Join(keyids, " "))
	}
	tw.Flush()
}
