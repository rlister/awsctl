package cmd

import (
	"bufio"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete NAME",
	Aliases: []string{"del"},
	Short:   "Delete secrets",
	Long:    `Delete secrets, optionally matching prefix.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 1:
			force, _ := cmd.Flags().GetBool("force")
			deleteSecret(args[0], force)
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func deleteSecret(name string, force bool) {
	input := bufio.NewScanner(os.Stdin)

	fmt.Printf("Delete %s? [y/n] ", name)
	input.Scan()

	if input.Text() == "y" {
		output, err := client.DeleteSecret(context.TODO(), &secretsmanager.DeleteSecretInput{SecretId: &name, ForceDeleteWithoutRecovery: &force})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Deletion date:", output.DeletionDate)
	}
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().BoolP("force", "f", false, "Force immediate delete")
}
