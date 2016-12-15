package cmd

import (
	"fmt"
	"os"

	"github.com/tonglil/labeler/utils"

	"github.com/spf13/cobra"
)

var (
	// Configuration options
	dryrun   bool
	token    string
	endpoint string
	repo     string

	// App info
	version bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "labeler",
	Short: "Manage labels on GitHub as code",
	Long: `
Labeler is a CLI application for managing labels on Github as code.

With the ability to scan and apply label changes, repository maintainers can now
empower contributors to submit PRs and improve the project management
process/label system!
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			fmt.Fprintf(os.Stdout, "version %s\n", utils.GetVersion())
			os.Exit(0)
		}

		cmd.Help()
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return RootCmd.Execute()
}

func init() {
	//cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.

	// Persistent flags, global for the application.
	RootCmd.PersistentFlags().BoolVarP(&dryrun, "dryrun", "d", false, "Show what would happen")

	RootCmd.PersistentFlags().StringVarP(&repo, "repo", "r", "", "GitHub repository (default is read from the file)")
	RootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "The GithHub token [overrides GITHUB_TOKEN]")
	RootCmd.PersistentFlags().StringVarP(&endpoint, "api", "a", utils.Api, "The GithHub API endpoint [overrides GITHUB_API]")

	// Local flags, only run when this action is called directly.
	RootCmd.Flags().BoolVarP(&version, "version", "V", false, "Show version")
}
