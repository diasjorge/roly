package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "roly",
	Short: "Easily work with AWS profiles",
	Long: `Set AWS environment variables based on your profile.
The goal is to make it easy to work with AWS profiles using assume role.
By working this way you don't need to have credentials in your accounts
if you can assume a role to it from a different account.`,
	SilenceUsage: true,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
