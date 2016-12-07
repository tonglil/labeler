package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang/glog"
	"github.com/tonglil/labeler/types"
	"github.com/tonglil/labeler/writer"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	api      = "https://api.github.com/"
	apiEnv   = "GITHUB_API"
	tokenEnv = "GITHUB_TOKEN"
)

var (
	// Deliberately uninitialized, see getVersion().
	version string
)

var (
	// Get labels
	scan bool

	// Configuration options
	dryrun   bool
	token    string
	endpoint string
	repo     string

	// App info
	help        bool
	versionFlag bool
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [<options>] <file.yaml>\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Manage labels on GitHub as code")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Available options:")

	flag.PrintDefaults()

	fmt.Fprintln(os.Stderr, "")

	os.Exit(1)
}

func init() {
	flag.BoolVar(&scan, "scan", false, "Scan the repo for label information")

	flag.BoolVar(&dryrun, "dry-run", false, "Show what would happen (default false)")
	flag.StringVar(&repo, "repo", "", "Use a different repository (default \"from file\")")
	flag.StringVar(&token, "token", "", "Use a different GithHub token [overrides GITHUB_TOKEN environment variable]")
	flag.StringVar(&endpoint, "endpoint", api, "Use a different GithHub API endpoint [overrides GITHUB_API environment variable]")

	flag.BoolVar(&help, "help", false, "Show help")
	flag.BoolVar(&versionFlag, "version", false, "Show version")

	flag.Usage = usage

	flag.Set("logtostderr", "true")
}

func main() {
	flag.Parse()

	if help || len(flag.Args()) != 1 {
		flag.Usage()
	}

	if versionFlag {
		fmt.Fprintf(os.Stdout, "version %s\n", getVersion())
		os.Exit(0)
	}

	file := flag.Args()[0]

	endpoint := getEndpoint(endpoint)

	token, err := getToken(token)
	if err != nil {
		fatal(err)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	err = setEndpoint(client, endpoint)
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
		//reader.Run(client, file, opt)
	} else {
		err = writer.Run(client, file, opt)
	}

	if err != nil {
		fatal(err)
	}

	os.Exit(0)
}
