package writer

import (
	"github.com/google/go-github/github"
	"github.com/tonglil/labeler/config"
	"github.com/tonglil/labeler/logs"
	"github.com/tonglil/labeler/remote"
	"github.com/tonglil/labeler/types"
)

// Run executes the write actions against the repo.
func Run(client *github.Client, opt *types.Options) error {
	file := opt.Filename

	lf, err := config.ReadFile(file)
	if err != nil {
		return err
	}

	opt.Repo, err = config.GetRepo(opt, lf)
	if err != nil {
		logs.V(0).Infof("No repo provided")
		return err
	}

	err = opt.ValidateRepo()
	if err != nil {
		logs.V(0).Infof("Failed to parse repo format: owner/name")
		return err
	}

	// Get all remote labels from repo
	labelsRemote, err := remote.GetLabels(client, opt)
	if err != nil {
		return err
	}

	var n, total int

	// Rename
	labels, n, err := Rename(client, opt, lf.Labels, labelsRemote)
	if err != nil {
		return err
	}

	logs.V(6).Infof("Finished renaming %d labels", n)
	total += n

	// Update
	labels, n, err = Update(client, opt, labels, labelsRemote)
	if err != nil {
		return err
	}

	logs.V(6).Infof("Finished updating %d labels", n)
	total += n

	// Create
	labels, n, err = Create(client, opt, labels, labelsRemote)
	if err != nil {
		return err
	}

	logs.V(6).Infof("Finished creating %d labels", n)
	total += n

	// Delete
	n, err = Delete(client, opt, lf.Labels, labelsRemote)
	if err != nil {
		return err
	}

	logs.V(6).Infof("Finished deleting %d labels", n)
	total += n

	logs.V(4).Infof("Processed %d labels in total", total)

	return nil
}
