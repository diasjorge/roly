package cmd

import (
	"errors"
	"fmt"

	"github.com/diasjorge/roly/credentials"
	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export PROFILE",
	Short: "Print export statements to use the profile",
	RunE:  export,
}

func init() {
	RootCmd.AddCommand(exportCmd)
}

func export(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("PROFILE required")
	}

	profileName := args[0]

	creds, err := credentials.Get(profileName, quiet)

	if err != nil {
		return err
	}

	fmt.Printf("export AWS_PROFILE='%s'\n", profileName)
	fmt.Printf("export AWS_ACCESS_KEY_ID='%s'\n", creds.AccessKeyID)
	fmt.Printf("export AWS_SECRET_ACCESS_KEY='%s'\n", creds.SecretAccessKey)

	if creds.SessionToken != "" {
		fmt.Printf("export AWS_SESSION_TOKEN='%s'\n", creds.SessionToken)
	}

	return nil
}
