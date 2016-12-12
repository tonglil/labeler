package main

import (
	"flag"
	"os"

	"github.com/tonglil/labeler/cmd"
	"github.com/tonglil/labeler/reader"
	"github.com/tonglil/labeler/types"
	"github.com/tonglil/labeler/utils"
	"github.com/tonglil/labeler/writer"

	"github.com/golang/glog"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	// Get labels
	scan bool

	// Configuration options
	dryrun   bool
	token    string
	endpoint string
	repo     string
)

func init() {
	//flag.BoolVar(&scan, "scan", false, "Scan the repo for label information")

	//flag.BoolVar(&dryrun, "dry-run", false, "Show what would happen (default false)")
	//flag.StringVar(&repo, "repo", "", "Use a different repository (default \"from file\")")
	//flag.StringVar(&token, "token", "", "Use a different GithHub token [overrides GITHUB_TOKEN environment variable]")
	//flag.StringVar(&endpoint, "endpoint", "", "Use a different GithHub API endpoint [overrides GITHUB_API environment variable]")

	//flag.Set("logtostderr", "true")
}

func main() {
	cmd.Execute()

	return

	flag.Parse()

	file := flag.Args()[0]

	endpoint := utils.GetEndpoint(endpoint)

	token, err := utils.GetToken(token)
	if err != nil {
		fatal(err)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	err = utils.SetEndpoint(client, endpoint)
	if err != nil {
		fatal(err)
	}

	opt := &types.Options{
		DryRun:   dryrun,
		Repo:     repo,
		Filename: file,
	}

	if opt.DryRun {
		glog.V(0).Infof("Dry run enabled - changes will not be applied")
	}

	if scan {
		err = reader.Run(client, file, opt)
	} else {
		err = writer.Run(client, file, opt)
	}

	if err != nil {
		fatal(err)
	}

	os.Exit(0)
}

func fatal(e error) {
	glog.V(0).Info(e)
	os.Exit(1)
}
