package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/oauth2"

	"github.com/google/go-github/github"
)

const (
	tokenEnv = "GITHUB_TOKEN"
)

var (
	version = "0.0.0"
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
	fmt.Fprintf(os.Stderr, "usage: %s [-version] [-help] [<options>] <file.yaml>\n", os.Args[0])
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

	flag.BoolVar(&dryrun, "dry-run", false, "Show what would happen")
	flag.StringVar(&token, "token", "", "Use a different GithHub token (default: \"from GITHUB_TOKEN environment variable\")")
	flag.StringVar(&endpoint, "endpoint", "https://api.github.com", "Use a different GithHub API endpoint")
	flag.StringVar(&repo, "repo", "", "Use a different repository (default: \"from file\")")

	flag.BoolVar(&help, "help", false, "Show help")
	flag.BoolVar(&versionFlag, "version", false, "Show version")

	flag.Usage = usage
}

// labeler labels.yaml
// labeler -scan -endpoint labels.yaml

func main() {
	flag.Parse()

	if help || len(flag.Args()) != 1 {
		flag.Usage()
	}

	if versionFlag {
		fmt.Fprintf(os.Stdout, "version %s\n", version)
		os.Exit(0)
	}

	if token == "" {
		token = os.Getenv(tokenEnv)
		if token == "" {
			fmt.Printf("missing environment variable %s\n", tokenEnv)
			os.Exit(1)
		}
	}

	fmt.Println("Starting...")

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)
	fmt.Println(client)

	os.Exit(0)
}
