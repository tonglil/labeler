package remote

import (
	"github.com/google/go-github/github"
	"github.com/tonglil/labeler/logs"
	"github.com/tonglil/labeler/types"
)

// GetLabels fetches all labels in a repository, iterating over pages for 50 at a time.
func GetLabels(client *github.Client, opt *types.Options) ([]*github.Label, error) {
	var labelsRemote []*github.Label

	pagination := &github.ListOptions{
		PerPage: 50,
		Page:    1,
	}

	for {
		logs.V(4).Infof("Fetching labels from Github, page %d", pagination.Page)

		labels, resp, err := client.Issues.ListLabels(opt.RepoOwner(), opt.RepoName(), pagination)
		if err != nil {
			logs.V(0).Infof("Failed to fetch labels from Github")
			return nil, err
		}
		logs.V(6).Infof("Response: %#v", resp)

		labelsRemote = append(labelsRemote, labels...)

		if resp.NextPage == 0 {
			logs.V(4).Infoln("Fetched all labels from Github")
			break
		}
		pagination.Page = resp.NextPage
	}

	return labelsRemote, nil
}
