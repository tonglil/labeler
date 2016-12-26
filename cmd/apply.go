package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tonglil/labeler/logs"
	"github.com/tonglil/labeler/types"
	"github.com/tonglil/labeler/utils"
	"github.com/tonglil/labeler/writer"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply file",
	Short: "Apply a YAML label definition file",
	Long: `
This command will apply the labels in "labels.yaml" to the "docker/docker"
repository without actually changing anything, just show what would be done:
  labeler apply labels.yaml -d -r docker/docker
	`,
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

		return writer.Run(client, opt)
	},
}

func init() {
	RootCmd.AddCommand(applyCmd)
}
