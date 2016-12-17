package writer

import (
	"github.com/google/go-github/github"
	"github.com/tonglil/labeler/logs"
	"github.com/tonglil/labeler/types"
)

func Rename(client *github.Client, opt *types.Options, local []*types.Label, remote []*github.Label) ([]*types.Label, int, error) {
	var remain []*types.Label
	var count int

	for _, l := range local {
		if l.From != "" {
			if _, ok := remoteHas(l.Name, remote); ok {
				logs.V(0).Infof("Skipped renaming '%s' to '%s', label already exists - please update your config file '%s'", l.From, l.Name, opt.Filename)
				continue
			}

			if r, ok := remoteHas(l.From, remote); ok {
				logs.V(4).Infof("Renaming '%s' to '%s' with color '%s' to '%s'", *r.Name, l.Name, *r.Color, l.Color)

				if opt.DryRun {
					count++
					continue
				}

				label, resp, err := client.Issues.EditLabel(opt.RepoOwner(), opt.RepoName(), *r.Name, &github.Label{
					Name:  &l.Name,
					Color: &l.Color,
				})
				if err != nil {
					logs.V(0).Infof("Failed to rename label '%s' to '%s' with color '%s' to '%s'", *r.Name, l.Name, *r.Color, l.Color)
					return nil, count, err
				}
				logs.V(6).Infof("Response: %s", resp)
				logs.V(4).Infof("Renamed label '%s'", label)

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
		if r, ok := remoteHas(l.Name, remote); ok && l.Color != *r.Color {
			logs.V(4).Infof("Updating '%s' with color '%s' to '%s'", l.Name, *r.Color, l.Color)

			if opt.DryRun {
				count++
				continue
			}

			label, resp, err := client.Issues.EditLabel(opt.RepoOwner(), opt.RepoName(), l.Name, &github.Label{
				Color: &l.Color,
			})
			if err != nil {
				logs.V(0).Infof("Failed to update label '%s' with color '%s' to '%s'", l.Name, *r.Color, l.Name)
				return nil, count, err
			}
			logs.V(6).Infof("Response: %s", resp)
			logs.V(4).Infof("Updated label '%s'", label)

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
			logs.V(4).Infof("Creating '%s' with color '%s'", l.Name, l.Color)

			if opt.DryRun {
				count++
				continue
			}

			label, resp, err := client.Issues.CreateLabel(opt.RepoOwner(), opt.RepoName(), &github.Label{
				Name:  &l.Name,
				Color: &l.Color,
			})
			if err != nil {
				logs.V(0).Infof("Failed to create label '%s' with color '%s'", l.Name, l.Color)
				return nil, count, err
			}
			logs.V(6).Infof("Response: %s", resp)
			logs.V(4).Infof("Created label '%s'", label)

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

		logs.V(4).Infof("Deleting '%s' with color '%s'", *l.Name, *l.Color)

		if opt.DryRun {
			count++
			continue
		}

		resp, err := client.Issues.DeleteLabel(opt.RepoOwner(), opt.RepoName(), *l.Name)
		if err != nil {
			logs.V(0).Infof("Failed to delete label '%s' with color '%s'", *l.Name, *l.Color)
			return count, err
		}
		logs.V(6).Infof("Response: %s", resp)
		logs.V(4).Infof("Deleted label '%s'", l)

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
