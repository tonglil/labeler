package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tonglil/versioning"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Run: func(cmd *cobra.Command, args []string) {
		versioning.Write(os.Stdout)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
