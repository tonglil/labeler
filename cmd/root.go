package cmd

import (
	"fmt"
	"os"

	"github.com/tonglil/labeler/logs"
	"github.com/tonglil/labeler/utils"

	"github.com/spf13/cobra"
)

var (
	// Configuration options
	dryrun   bool
	token    string
	endpoint string
	repo     string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "labeler",
	Short: "Manage labels on GitHub as code",
	Long: `
Labeler is a CLI application for managing labels on Github as code.

With the ability to scan and apply label changes, repository maintainers can
empower contributors to submit PRs and improve the project management
process/label system!
	`,
	SilenceErrors: true,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

// Define your flags and configuration settings.
func init() {
	// Persistent flags, global for the application.
	RootCmd.PersistentFlags().BoolVarP(&dryrun, "dryrun", "d", false, "Show what would happen")

	RootCmd.PersistentFlags().StringVarP(&repo, "repo", "r", "", "GitHub repository (default is read from the file)")
	RootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "The GithHub token [overrides GITHUB_TOKEN]")
	RootCmd.PersistentFlags().StringVarP(&endpoint, "api", "a", utils.Api, "The GithHub API endpoint [overrides GITHUB_API]")

	RootCmd.PersistentFlags().IntVarP(&logs.Threshold, "level", "l", 1, "The maximum level of logging to display")
}
