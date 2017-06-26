package cmd

import (
	"os"
	osExec "os/exec"

	"github.com/diasjorge/roly/credentials"
	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec PROFILE CMD",
	Short: "Execute command with AWS environment variables set",
	RunE:  exec,
}

func init() {
	RootCmd.AddCommand(execCmd)
}

func exec(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return cmd.Usage()
	}

	creds, err := credentials.Get(args[0], quiet)

	if err != nil {
		return err
	}

	os.Setenv("AWS_ACCESS_KEY_ID", creds.AccessKeyID)
	os.Setenv("AWS_SECRET_ACCESS_KEY", creds.SecretAccessKey)
	os.Setenv("AWS_SESSION_TOKEN", creds.SessionToken)

	subCommandName := args[1]
	subCommandArgs := args[2:]

	subCommand := osExec.Command(subCommandName, subCommandArgs...)
	subCommand.Stdout = os.Stdout

	if err := subCommand.Run(); err != nil {
		return err
	}

	return nil
}
