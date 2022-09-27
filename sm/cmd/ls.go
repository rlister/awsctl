package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls [PREFIX]",
	Short: "List secrets",
	Long:  `List secrets, optionally matching prefix.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			listSecrets("")
		case 1:
			listSecrets(args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}

func listSecrets(prefix string) {
	for _, s := range secrets(prefix) {
		fmt.Println(*s.Name)
	}
}
