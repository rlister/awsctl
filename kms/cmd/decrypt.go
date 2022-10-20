package cmd

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/spf13/cobra"
	"log"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:     "decrypt TEXT",
	Aliases: []string{"dec"},
	Short:   "Decrypt text",
	Long:    `Decrypt given text blob.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 1:
			decryptKmsKey(args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func decryptKmsKey(text string) {
	blob, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		log.Fatal("Base64 error:", err)
	}

	output, err := client.Decrypt(context.TODO(), &kms.DecryptInput{CiphertextBlob: blob})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("KeyId:", *output.KeyId)
	fmt.Printf("Plaintext: %s\n", string(output.Plaintext))
}

func init() {
	rootCmd.AddCommand(decryptCmd)
}
