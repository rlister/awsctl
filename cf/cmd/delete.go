package cmd

import (
	"bufio"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete PREFIX",
	Aliases: []string{"del"},
	Short:   "delete stacks",
	Long:    `Delete stacks matching given PREFIX. Prompts for confirmation.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("wrong number of arguments")
		}
		deleteStacks(args[0])
	},
}

// deleteStacks with given prefix
func deleteStacks(prefix string) {
	input := bufio.NewScanner(os.Stdin)

	for _, s := range stacks(prefix) {
		fmt.Printf("Delete %s? [y/n] ", *s.StackName)
		input.Scan()

		if input.Text() == "y" {
			_, err := client.DeleteStack(context.TODO(), &cloudformation.DeleteStackInput{StackName: s.StackName})
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
