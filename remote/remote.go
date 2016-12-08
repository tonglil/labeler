package remote

import (
	"github.com/tonglil/labeler/types"

	"github.com/golang/glog"
	"github.com/google/go-github/github"
)

// GetLabels fetches all labels in a repository, iterating over pages for 50 at a time.
func GetLabels(client *github.Client, opt *types.Options) ([]*github.Label, error) {
	var labelsRemote []*github.Label

	pagination := &github.ListOptions{
		PerPage: 50,
		Page:    1,
	}

	for {
		glog.V(4).Infof("Fetching labels from Github, page %d", pagination.Page)

		labels, resp, err := client.Issues.ListLabels(opt.RepoOwner(), opt.RepoName(), pagination)
		if err != nil {
			glog.V(0).Infof("Failed to fetch labels from Github")
			return nil, err
		}
		glog.V(6).Infof("Response: %s", resp)

		labelsRemote = append(labelsRemote, labels...)

		if resp.NextPage == 0 {
			glog.V(4).Info("Fetched all labels from Github")
			break
		}
		pagination.Page = resp.NextPage
	}

	return labelsRemote, nil
}
