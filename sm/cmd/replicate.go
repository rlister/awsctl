package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/spf13/cobra"
	"log"
)

// replicateCmd represents the replicate command
var replicateCmd = &cobra.Command{
	Use:     "replicate NAME REGION [REGION ...]",
	Aliases: []string{"rep", "repl"},
	Short:   "Replicate secrets",
	Long:    `Replicate secrets, optionally matching prefix.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			log.Fatal("wrong number of arguments")
		case 1:
			log.Fatal("add regions to replicate")
		default:
			replicateSecret(args)
		}
	},
}

// replicate secret to list of regions
func replicateSecret(args []string) {
	output, err := client.ReplicateSecretToRegions(context.TODO(), &secretsmanager.ReplicateSecretToRegionsInput{
		SecretId:          &args[0],
		AddReplicaRegions: regions(args[1:]),
	})

	if err != nil {
		log.Fatal(err)
	}

	for _, s := range output.ReplicationStatus {
		fmt.Println(*s.Region)
	}
}

// take array of region strings and return array of ReplicaRegionType
func regions(names []string) []types.ReplicaRegionType {
	regions := []types.ReplicaRegionType{}
	for _, name := range names {
		regions = append(regions, types.ReplicaRegionType{Region: &name})
	}
	return regions
}

func init() {
	rootCmd.AddCommand(replicateCmd)
}
