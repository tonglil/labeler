package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tonglil/labeler/config"
	"github.com/tonglil/labeler/logs"
	"github.com/tonglil/labeler/types"
)

var trimCmd = &cobra.Command{
	Use:          "trim file",
	Short:        "Trim the YAML definition file of all 'from:'s",
	Long:         `Remove 'from:'s from a file after they have been applied`,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			fmt.Println(cmd.UsageString())
			return fmt.Errorf("no file given")
		}

		file := args[0]

		lf, err := config.ReadFile(file)
		if err != nil {
			return err
		}

		var count int
		var labels []*types.Label

		for _, label := range lf.Labels {
			if label.From != "" {
				logs.V(4).Infof("Removing 'from: %s' for label '%s'", label.From, label.Name)
				label.From = ""
				count++
			}
			labels = append(labels, label)
		}

		lf.Labels = labels

		err = config.WriteFile(file, lf)
		if err != nil {
			return err
		}

		logs.V(4).Infof("Trimmed %d labels in total", count)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(trimCmd)

	trimCmd.PersistentFlags().IntVarP(&logs.Threshold, "level", "l", 1, "The maximum level of logging to display")
}
