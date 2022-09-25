package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
)

const dateFormat = "2006-01-02 15:04 MST"

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get [REPO]",
	Aliases: []string{"g"},
	Short:   "Get details for repos or images",
	Long:    `Get details for repos, or images in repo if repo name given.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			getRepos()
		case 1:
			getImages(args[0])
		default:
			log.Fatal("wrong number of arguments")
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func getRepos() {
	const format = "%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "NAME", "CREATED", "URI")
	for _, r := range repos() {
		fmt.Fprintf(tw, format, *r.RepositoryName, (*r.CreatedAt).Format(dateFormat), *r.RepositoryUri)
	}
	tw.Flush()
}

func getImages(repo string) {
	const format = "%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "PUSHED", "SIZE(MB)", "TAGS")
	for _, i := range imageDetails(repo) {
		sort.Strings(i.ImageTags)
		tags := strings.Join(i.ImageTags, " ")
		fmt.Fprintf(tw, format, (*i.ImagePushedAt).Format(dateFormat), (*i.ImageSizeInBytes)/1000000, tags)
	}
	tw.Flush()
}
