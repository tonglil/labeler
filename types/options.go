package types

import (
	"fmt"
	"strings"
)

// Options passed during the execution of Labeler.
type Options struct {
	DryRun   bool
	Repo     string
	Filename string
	owner    string
	name     string
}

// ValidateRepo checks the repo name is well formatted.
func (o *Options) ValidateRepo() error {
	_, _, err := split(o.Repo)
	if err != nil {
		return err
	}

	return nil
}

// RepoOwner gets the repo's owner from the repo string.
func (o *Options) RepoOwner() string {
	if o.owner != "" {
		return o.owner
	}

	o.owner, _, _ = split(o.Repo)

	return o.owner
}

// RepoName gets the repo's name from the repo string.
func (o *Options) RepoName() string {
	if o.name != "" {
		return o.name
	}

	_, o.name, _ = split(o.Repo)

	return o.name
}

func split(repo string) (string, string, error) {
	p := strings.Split(repo, "/")
	if len(p) != 2 {
		return "", "", fmt.Errorf("malformed repo format")
	}

	return p[0], p[1], nil
}
