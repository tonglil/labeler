package writer

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/tonglil/labeler/types"

	"github.com/golang/glog"
	"github.com/google/go-github/github"
	yaml "gopkg.in/yaml.v2"
)

// Run executes the write actions against the repo.
func Run(client *github.Client, file string, opt *types.Options) error {
	lf, err := ReadConfigFile(file)
	if err != nil {
		return err
	}

	opt.Repo, err = GetRepo(opt, lf)
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
	labelsRemote, err := GetRemoteLabels(client, opt)
	if err != nil {
		return err
	}

	var n, total int

	// Rename
	labels, n, err := Rename(client, opt, lf.Labels, labelsRemote)
	if err != nil {
		return err
	}

	glog.V(6).Infof("Finished renaming %d labels", n)
	total += n

	// Update
	labels, n, err = Update(client, opt, labels, labelsRemote)
	if err != nil {
		return err
	}

	glog.V(6).Infof("Finished updating %d labels", n)
	total += n

	// Create
	labels, n, err = Create(client, opt, labels, labelsRemote)
	if err != nil {
		return err
	}

	glog.V(6).Infof("Finished creating %d labels", n)
	total += n

	// Delete
	n, err = Delete(client, opt, lf.Labels, labelsRemote)
	if err != nil {
		return err
	}

	glog.V(6).Infof("Finished deleting %d labels", n)
	total += n

	glog.V(6).Infof("Processed %d labels in total", total)

	return nil
}

func Rename(client *github.Client, opt *types.Options, local []*types.Label, remote []*github.Label) ([]*types.Label, int, error) {
	var remain []*types.Label
	var count int

	for _, l := range local {
		if l.From != "" {
			if _, ok := remoteHas(l.Name, remote); ok {
				glog.Infof("Skipped renaming '%s' to '%s', label already exists - please update your config file '%s'", l.From, l.Name, opt.Filename)
				continue
			}

			if r, ok := remoteHas(l.From, remote); ok {
				count++
				glog.V(4).Infof("Renaming '%s' to '%s' with color '%s' to '%s'\n", *r.Name, l.Name, *r.Color, l.Color)
				continue
			}
		}

		remain = append(remain, l)
	}

	return remain, count, nil
}

func Update(client *github.Client, opt *types.Options, local []*types.Label, remote []*github.Label) ([]*types.Label, int, error) {
	var remain []*types.Label
	var count int

	for _, l := range local {
		if r, ok := remoteHas(l.Name, remote); ok {
			count++
			glog.V(4).Infof("Updating '%s' with color '%s' to '%s'\n", l.Name, *r.Color, l.Color)
			continue
		}

		remain = append(remain, l)
	}

	return remain, count, nil
}

func Create(client *github.Client, opt *types.Options, local []*types.Label, remote []*github.Label) ([]*types.Label, int, error) {
	var remain []*types.Label
	var count int

	for _, l := range local {
		if _, ok := remoteHas(l.Name, remote); !ok {
			count++
			glog.V(4).Infof("Creating '%s' with color '%s'\n", l.Name, l.Color)
			continue
		}

		remain = append(remain, l)
	}

	return remain, count, nil
}

func Delete(client *github.Client, opt *types.Options, local []*types.Label, remote []*github.Label) (int, error) {
	var count int

	for _, l := range remote {
		if _, ok := localHasOrRenamed(*l.Name, local); ok {
			continue
		}

		count++
		glog.V(4).Infof("Deleting '%s' with color '%s'\n", *l.Name, *l.Color)
	}

	return count, nil
}

func remoteHas(name string, labels []*github.Label) (*github.Label, bool) {
	for _, l := range labels {
		if name == *l.Name {
			return l, true
		}
	}

	return nil, false
}

func localHasOrRenamed(name string, labels []*types.Label) (*types.Label, bool) {
	for _, l := range labels {
		if name == l.Name || name == l.From {
			return l, true
		}
	}

	return nil, false
}

// GetRemoteLabels fetches all labels in a repository, iterating over pages for 50 at a time.
func GetRemoteLabels(client *github.Client, opt *types.Options) ([]*github.Label, error) {
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
		glog.V(4).Infof("Response: %s", resp)

		labelsRemote = append(labelsRemote, labels...)

		if resp.NextPage == 0 {
			glog.V(4).Info("Fetched all labels from Github")
			break
		}
		pagination.Page = resp.NextPage
	}

	return labelsRemote, nil
}

// GetRepo configures the repo being used as determined by the option, and then the label file.
func GetRepo(opt *types.Options, lf *types.LabelFile) (string, error) {
	if opt.Repo != "" {
		return opt.Repo, nil
	}

	if lf.Repo != "" {
		return lf.Repo, nil
	}

	return "", fmt.Errorf("no repo")
}

// ReadConfigFile opens the label file and reads its contents into a LabelFile.
func ReadConfigFile(file string) (*types.LabelFile, error) {
	path, err := filepath.Abs(file)
	if err != nil {
		glog.V(0).Infof("Failed to find %s", file)
		return nil, err
	}

	f, err := os.Open(path)
	if err != nil {
		glog.V(0).Infof("Failed to open %s", path)
		return nil, err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		glog.V(0).Infof("Failed to read %s", path)
		return nil, err
	}

	glog.V(4).Infof("Read file %s", path)

	lf := types.LabelFile{}

	err = yaml.Unmarshal(data, &lf)
	if err != nil {
		glog.V(0).Infof("Failed to unmarshal %s", path)
		return nil, err
	}

	return &lf, nil
}
