package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls [REPO]",
	Short: "List repos or images",
	Long:  `List repos, or images in repo if repo name given.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			listRepos()
		case 1:
			listImages(args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}

func listRepos() {
	for _, r := range repos() {
		fmt.Println(*r.RepositoryName)
	}
}

func listImages(repo string) {
	for _, i := range images(repo) {
		if i.ImageTag != nil {
			fmt.Println(*i.ImageTag)
		}
	}
}
