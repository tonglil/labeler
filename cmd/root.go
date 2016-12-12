package cmd

import (
	"fmt"
	"os"

	"github.com/tonglil/labeler/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	api = "https://api.github.com/"
	//apiEnv   = "GITHUB_API"
	//tokenEnv = "GITHUB_TOKEN"
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
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if version {
		fmt.Fprintf(os.Stdout, "version %s\n", utils.GetVersion())
		os.Exit(0)
	}
}

func init() {
	//cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.

	// Persistent flags, global for the application.
	RootCmd.PersistentFlags().BoolVarP(&dryrun, "dryrun", "d", false, "Show what would happen")

	RootCmd.PersistentFlags().StringVarP(&repo, "repo", "r", "", "GitHub repository (default is read from the file)")
	RootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "The GithHub token [overrides GITHUB_TOKEN]")
	RootCmd.PersistentFlags().StringVarP(&endpoint, "api", "a", api, "The GithHub API endpoint [overrides GITHUB_API]")

	// Local flags, only run when this action is called directly.
	RootCmd.Flags().BoolVarP(&version, "version", "v", false, "Show version")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	//if cfgFile != "" { // enable ability to specify config file via flag
	//viper.SetConfigFile(cfgFile)
	//}

	//viper.SetConfigName(".cobra") // name of config file (without extension)
	//viper.AddConfigPath("$HOME")  // adding home directory as first search path
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
