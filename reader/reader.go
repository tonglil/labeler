package reader

import (
	"encoding/json"

	"github.com/tonglil/labeler/config"
	"github.com/tonglil/labeler/remote"
	"github.com/tonglil/labeler/types"

	"github.com/golang/glog"
	"github.com/google/go-github/github"
	yaml "gopkg.in/yaml.v2"
)

// Run executes the write actions against the repo.
func Run(client *github.Client, file string, opt *types.Options) error {
	// TODO:
	// DryRun should cleanup if missing as well...
	err := config.CreateIfMissing(file)
	if err != nil {
		return err
	}

	lf, err := config.ReadFile(file)
	if err != nil {
		return err
	}

	opt.Repo, err = config.GetRepo(opt, lf)
	if err != nil {
		glog.V(0).Infof("No repo provided")
		return err
	}

	err = opt.ValidateRepo()
	if err != nil {
		glog.V(0).Infof("Failed to parse repo format: owner/name")
		return err
	}

	// Get all remote labels from repo
	labelsRemote, err := remote.GetLabels(client, opt)
	if err != nil {
		return err
	}

	total := len(labelsRemote)

	x, err := json.Marshal(labelsRemote)
	if err != nil {
		glog.V(0).Infof("Failed to marshal labels from remote format")
		return err
	}

	labels := []*types.Label{}

	// TODO:
	// Can we directly unmarshal from labelsRemote?
	err = yaml.Unmarshal(x, &labels)
	if err != nil {
		glog.V(0).Infof("Failed to unmarshal labels to local format")
		return err
	}

	for _, l := range labels {
		glog.V(4).Infof("Fetched '%s' with color '%s'\n", l.Name, l.Color)
	}

	lf = &types.LabelFile{
		Repo:   opt.Repo,
		Labels: labels,
	}

	if !opt.DryRun {
		err = config.WriteFile(file, lf)
		if err != nil {
			return err
		}
	}

	glog.V(4).Infof("Processed %d labels in total", total)

	return nil
}
