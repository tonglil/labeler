package writer

import (
	"context"

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
				logs.V(4).Infof("Renaming '%s' to '%s' with color '%s' to '%s'", r.Name, l.Name, r.Color, l.Color)

				if opt.DryRun {
					count++
					continue
				}

				label, resp, err := client.Issues.EditLabel(context.Background(), opt.RepoOwner(), opt.RepoName(), r.Name, toRemote(l))
				if err != nil {
					logs.V(0).Infof("Failed to rename label '%s' to '%s' with color '%s' to '%s'", r.Name, l.Name, r.Color, l.Color)
					return nil, count, err
				}
				logs.V(6).Infof("Response: %#v", resp)
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
		r, ok := remoteHas(l.Name, remote)
		if !ok || (l.Color == r.Color && l.Description == r.Description) {
			remain = append(remain, l)
			continue
		}

		if l.Color != r.Color {
			logs.V(4).Infof("Updating '%s' with color '%s' to '%s'", l.Name, r.Color, l.Color)
		}

		if l.Description != r.Description {
			logs.V(4).Infof("Updating '%s' with description '%s' to '%s'", l.Name, r.Description, l.Description)
		}

		if opt.DryRun {
			count++
			continue
		}

		label, resp, err := client.Issues.EditLabel(context.Background(), opt.RepoOwner(), opt.RepoName(), l.Name, toRemote(l))
		if err != nil {
			logs.V(0).Infof("Failed to update label '%s' with color '%s' to '%s' and description '%s' to '%s'", l.Name, r.Color, l.Color, r.Description, l.Description)
			return nil, count, err
		}
		logs.V(6).Infof("Response: %#v", resp)
		logs.V(4).Infof("Updated label '%s'", label)

		count++
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

			label, resp, err := client.Issues.CreateLabel(context.Background(), opt.RepoOwner(), opt.RepoName(), &github.Label{
				Name:        &l.Name,
				Color:       &l.Color,
				Description: &l.Description,
			})
			if err != nil {
				logs.V(0).Infof("Failed to create label '%s' with color '%s'", l.Name, l.Color)
				return nil, count, err
			}
			logs.V(6).Infof("Response: %#v", resp)
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

		resp, err := client.Issues.DeleteLabel(context.Background(), opt.RepoOwner(), opt.RepoName(), *l.Name)
		if err != nil {
			logs.V(0).Infof("Failed to delete label '%s' with color '%s'", *l.Name, *l.Color)
			return count, err
		}
		logs.V(6).Infof("Response: %#v", resp)
		logs.V(4).Infof("Deleted label '%s'", l)

		count++
	}

	return count, nil
}

func remoteHas(name string, labels []*github.Label) (*types.Label, bool) {
	for _, l := range labels {
		if name == *l.Name {
			return toLocal(l), true
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

func toLocal(r *github.Label) *types.Label {
	if r == nil {
		return nil
	}

	l := &types.Label{}
	if r.Name != nil {
		l.Name = *r.Name
	}

	if r.Color != nil {
		l.Color = *r.Color
	}

	if r.Description != nil {
		l.Description = *r.Description
	}

	return l
}

func toRemote(l *types.Label) *github.Label {
	if l == nil {
		return nil
	}

	return &github.Label{
		Name:        &l.Name,
		Color:       &l.Color,
		Description: &l.Description,
	}
}
