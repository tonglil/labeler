package cmd

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/tonglil/labeler/types"
	"github.com/tonglil/labeler/utils"
	"github.com/tonglil/labeler/writer"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply file",
	Short: "Apply a YAML label definition file",
	Long: `
A longer description that spans multiple lines and likely contains examples to
quickly create a Cobra application.
	`,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
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
			glog.V(0).Infof("Dry run enabled - changes will not be applied")
		}

		return writer.Run(client, opt)
	},
}

func init() {
	RootCmd.AddCommand(applyCmd)
}
