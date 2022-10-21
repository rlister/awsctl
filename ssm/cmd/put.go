package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/spf13/cobra"
	"log"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:     "put",
	Aliases: []string{"p"},
	Short:   "Put parameter",
	Long:    `Put parameter.`,
	Run: func(cmd *cobra.Command, args []string) {
		overwrite, _ := cmd.Flags().GetBool("overwrite")

		output, err := client.PutParameter(context.TODO(), &ssm.PutParameterInput{
			Name:      &args[0],
			Value:     &args[1],
			Type:      "String",
			Overwrite: &overwrite,
		})

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("version:", output.Version)
	},
}

func putParam(name, value, typ *string) {
	ovw := false

	output, err := client.PutParameter(context.TODO(), &ssm.PutParameterInput{
		Name:      name,
		Value:     value,
		Type:      types.ParameterType(*typ),
		Overwrite: &ovw,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("version:", output.Version)
}

func init() {
	rootCmd.AddCommand(putCmd)
	putCmd.Flags().BoolP("overwrite", "o", false, "Overwrite existing parameter")
}
