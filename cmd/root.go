package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
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
