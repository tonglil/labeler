package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tonglil/labeler/logs"
	"github.com/tonglil/labeler/reader"
	"github.com/tonglil/labeler/types"
	"github.com/tonglil/labeler/utils"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan file",
	Short: "Save a repository's labels into a YAML definition file",
	Long: `Save remote labels into a file

Example:
  $ labeler scan labels.yaml -r docker/docker -l 5

  Scan the labels from the "docker/docker" repository into a
  file called "labels.yaml", logging what happened.`,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			fmt.Println(cmd.UsageString())
			return fmt.Errorf("no file given")
		}

		file := args[0]
		client, err := utils.GetClient(endpoint, token)
		if err != nil {
			return err
		}

		opt := &types.Options{
			DryRun:   cmd.Flag("dryrun").Changed,
			Repo:     cmd.Flag("repo").Value.String(),
			Filename: file,
		}

		if opt.DryRun {
			logs.V(0).Infof("Dry run enabled - changes will not be applied")
		}

		return reader.Run(client, opt)
	},
}

func init() {
	RootCmd.AddCommand(scanCmd)

	scanCmd.PersistentFlags().BoolVarP(&dryrun, "dryrun", "d", false, "Show what would happen")

	scanCmd.PersistentFlags().StringVarP(&repo, "repo", "r", "", "GitHub repository (default is read from the file)")
	scanCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "The GithHub token [overrides GITHUB_TOKEN]")
	scanCmd.PersistentFlags().StringVarP(&endpoint, "api", "a", utils.Api, "The GithHub API endpoint [overrides GITHUB_API]")

	scanCmd.PersistentFlags().IntVarP(&logs.Threshold, "level", "l", 1, "The maximum level of logging to display")
}
