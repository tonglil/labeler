package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Output shell completion code for tab completion",
	Long: `Print shell code for evaluation of interactive completion of labeler commands

Examples:
  $ source <(labeler completion)

  Note that this depends on the bash-completion framework. It must be sourced
  before sourcing the labeler completion, e.g. on the Mac:

  $ brew install bash-completion
  $ source $(brew --prefix)/etc/bash_completion
  $ source <(labeler completion)`,
	Run: func(cmd *cobra.Command, args []string) {
		RootCmd.GenBashCompletion(os.Stdout)
	},
}

func init() {
	RootCmd.AddCommand(completionCmd)
}
