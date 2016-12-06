package writer

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"

	"github.com/tonglil/labeler/types"

	"github.com/golang/glog"
	"github.com/google/go-github/github"
)

var steps = `
1. Rename: intersection of local (l.From) + remote (l.Name)
2. Update: intersection of local (l.Name) + remote (l.Name)
3. Create: l.Name compliment of remote
4. Delete: l.Name compliment of locale
`

func Run(client *github.Client, file string, opt *types.Options) error {
	// open file
	// unmarshal file contents
	// configure the right repo
	// ep = [https://api.github.com/] repos/ [tonglil/labeler] /labels
	// get all remote labels from repo
	// for l in local labels with "from:"
	// if l doesn't exist
	//   if l.From exists
	//     rename from to it
	//       PATCH ep + /:name <- l.from
	//         {
	//           "name": "l.name",
	//           "color": "l.color"
	//         }
	//   else create it
	//     POST ep
	//       {
	//         "name": "l.name",
	//         "color": "l.color"
	//       }
	// do create non-existing local labels
	// do update existing local labels
	// do delete remaining remote labels
	lf, err := ReadConfigFile(file)
	if err != nil {
		return err
	}

	if opt.Repo == "" {
		opt.Repo = lf.Repo
	}

	if opt.Repo == "" {
		glog.V(0).Infof("No repo provided")
		return fmt.Errorf("no repo")
	}

	return nil
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
