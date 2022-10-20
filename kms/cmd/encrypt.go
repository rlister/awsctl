package cmd

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/spf13/cobra"
	"log"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:     "encrypt TEXT",
	Aliases: []string{"enc"},
	Short:   "Encrypt text",
	Long:    `Encrypt given text.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 2:
			encryptKmsKey(args[0], args[1])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func encryptKmsKey(key, text string) {
	output, err := client.Encrypt(context.TODO(), &kms.EncryptInput{KeyId: &key, Plaintext: []byte(text)})
	if err != nil {
		log.Fatal(err)
	}

	blob := base64.StdEncoding.EncodeToString(output.CiphertextBlob)
	fmt.Println(blob)
}

func init() {
	rootCmd.AddCommand(encryptCmd)
}
