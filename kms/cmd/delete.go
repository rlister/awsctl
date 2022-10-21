package cmd

import (
	"bufio"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete KEY",
	Aliases: []string{"del"},
	Short:   "Delete key",
	Long:    `Delete key by id.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 1:
			deleteKey(args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func deleteKey(id string) {
	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("Delete %s? [y/n] ", id)
	input.Scan()

	if input.Text() == "y" {
		output, err := client.ScheduleKeyDeletion(context.TODO(), &kms.ScheduleKeyDeletionInput{KeyId: &id})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("DeletionDate:", *output.DeletionDate)
	}
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
