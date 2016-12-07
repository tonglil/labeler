package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/tonglil/labeler/types"

	"github.com/golang/glog"
	yaml "gopkg.in/yaml.v2"
)

// ReadFile opens the label file and reads its contents into a LabelFile.
func ReadFile(file string) (*types.LabelFile, error) {
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
