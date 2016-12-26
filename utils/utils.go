package utils

import (
	"fmt"
	"net/url"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	// The default GitHub API url
	Api      = "https://api.github.com/"
	apiEnv   = "GITHUB_API"
	tokenEnv = "GITHUB_TOKEN"
)

var (
	// Deliberately uninitialized, see GetVersion().
	version string
)

func GetClient(endpoint string, token string) (*github.Client, error) {
	endpoint = getEndpoint(endpoint)

	token, err := getToken(token)
	if err != nil {
		return nil, err
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: token,
		},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	err = setEndpoint(client, endpoint)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func getToken(t string) (string, error) {
	// If token is set, use that.
	if t != "" {
		return t, nil
	}

	// If token is set as an environment variable, use that.
	t = os.Getenv(tokenEnv)
	if t != "" {
		return t, nil
	}

	// Otherwise return an error.
	return "", fmt.Errorf("missing environment variable %s", tokenEnv)
}

func getEndpoint(e string) string {
	// If endpoint is different from the default, use that.
	if e != Api {
		return e
	}

	// If endpoint is set as an environment variable, use that.
	e = os.Getenv(apiEnv)
	if e != "" {
		return e
	}

	// Otherwise use the default endpoint.
	return Api
}

func setEndpoint(c *github.Client, e string) error {
	if e != Api {
		ep, err := url.Parse(e)
		if err != nil {
			return err
		}

		c.BaseURL = ep
	}

	return nil
}

func GetVersion() string {
	if version != "" {
		return version
	}

	return "unknown"
}
