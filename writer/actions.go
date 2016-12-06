package writer

import (
	"github.com/tonglil/labeler/types"

	"github.com/golang/glog"
	"github.com/google/go-github/github"
)

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
				glog.V(4).Infof("Renaming '%s' to '%s' with color '%s' to '%s'\n", *r.Name, l.Name, *r.Color, l.Color)
				count++
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
			glog.V(4).Infof("Updating '%s' with color '%s' to '%s'\n", l.Name, *r.Color, l.Color)
			count++
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
			glog.V(4).Infof("Creating '%s' with color '%s'\n", l.Name, l.Color)
			count++
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

		glog.V(4).Infof("Deleting '%s' with color '%s'\n", *l.Name, *l.Color)
		count++
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
