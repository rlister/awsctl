package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:     "create NAME",
	Aliases: []string{"cre"},
	Short:   "Create secret",
	Long:    `Create secret with given name.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 1:
			createSecret(args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

// create a new secret from edited json string
func createSecret(name string) {
	data := edit("{}")

	output, err := client.CreateSecret(context.TODO(), &secretsmanager.CreateSecretInput{Name: &name, SecretString: &data})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(*output.ARN)
}

// get user editor from env
func editor() string {
	ed, ok := os.LookupEnv("EDITOR")
	if !ok {
		log.Fatal("please set environment variable EDITOR")
	}
	return ed
}

// edit data in user EDITOR, and return new string
func edit(data string) string {
	f, err := os.CreateTemp("", "awsctl-*.json")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	defer os.Remove(f.Name())

	// write current value to tmpfile
	if _, err := f.Write([]byte(data)); err != nil {
		log.Fatal(err)
	}

	// run editor on tmpfile
	cmd := exec.Command(editor(), f.Name())
	if err = cmd.Run(); err != nil {
		log.Fatal(err)
	}

	// read back edited content
	edited, err := ioutil.ReadFile(f.Name())
	if err != nil {
		log.Fatal(err)
	}

	return string(edited)
}

func init() {
	rootCmd.AddCommand(createCmd)
}
