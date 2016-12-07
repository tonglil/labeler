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

func ReadFileOrCreate(file string) (*types.LabelFile, error) {
	path, err := filepath.Abs(file)
	if err != nil {
		glog.V(0).Infof("Failed to find %s", file)
		return nil, err
	}

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		glog.V(0).Infof("Creating file %s", path)

		f, err := os.Create(file)
		if err != nil {
			glog.V(0).Infof("Failed to create file %s", file)
			return nil, err
		}
		f.Close()
	}

	return ReadFile(file)
}

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

// WriteFile opens the label file and overwrites the LabelFile into its contents.
func WriteFile(file string, lf *types.LabelFile) error {
	path, err := filepath.Abs(file)
	if err != nil {
		glog.V(0).Infof("Failed to find %s", file)
		return err
	}

	data, err := yaml.Marshal(lf)
	if err != nil {
		glog.V(0).Infof("Failed to marshal %T", lf)
		return err
	}

	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		glog.V(0).Infof("Failed to write %s", path)
		return err
	}

	glog.V(4).Infof("Wrote file %s", path)

	return nil
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
