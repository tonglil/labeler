package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/google/go-github/github"
)

func fatal(e error) {
	fmt.Fprintln(os.Stderr, e)
	os.Exit(1)
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
	if e != api {
		return e
	}

	// If endpoint is set as an environment variable, use that.
	e = os.Getenv(apiEnv)
	if e != "" {
		return e
	}

	// Otherwise use the default endpoint.
	return api
}

func setEndpoint(c *github.Client, e string) error {
	if e != api {
		ep, err := url.Parse(e)
		if err != nil {
			return err
		}

		c.BaseURL = ep
	}

	return nil
}

func getVersion() string {
	if version != "" {
		return version
	}

	return "unknown"
}
