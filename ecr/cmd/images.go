package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"log"
	"sort"
)

// get array of ImageIdentifiers for repo
func images(repo string) []types.ImageIdentifier {
	paginator := ecr.NewListImagesPaginator(client, &ecr.ListImagesInput{RepositoryName: &repo})

	images := []types.ImageIdentifier{}

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		for _, i := range output.ImageIds {
			images = append(images, i)
		}
	}

	return images
}

// sort arrays of ImageDetail
type byPushedAt []types.ImageDetail

func (x byPushedAt) Len() int           { return len(x) }
func (x byPushedAt) Less(i, j int) bool { return (*x[i].ImagePushedAt).After(*x[j].ImagePushedAt) }
func (x byPushedAt) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// return array of ImageDetails, sorted by most recent pushed
func imageDetails(repo string) []types.ImageDetail {
	paginator := ecr.NewDescribeImagesPaginator(client, &ecr.DescribeImagesInput{RepositoryName: &repo})

	images := []types.ImageDetail{}

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		for _, i := range output.ImageDetails {
			images = append(images, i)
		}
	}
	sort.Sort(byPushedAt(images))
	return images
}
