package cmd

import (
	"os"
	osExec "os/exec"
	"strings"

	"github.com/diasjorge/roly/credentials"
	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec PROFILE CMD",
	Short: "Execute command with AWS environment variables set",
	Long: `Execute command with AWS environment variables set

If you need to pass flags to your command, you need to specify where
the flags for the roly command finish using "--" like:
roly exec -q -- PROFILE CMD --your-command-flags
or you can quote your command like:
roly exec PROFILE "CMD --your-command-flags"`,
	RunE: exec}

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

	var subCommandName string
	var subCommandArgs []string

	// Assume command is quoted
	if len(args) == 2 {
		fields := strings.Fields(args[1])
		subCommandName = fields[0]
		subCommandArgs = fields[1:]
	} else {
		subCommandName = args[1]
		subCommandArgs = args[2:]
	}

	subCommand := osExec.Command(subCommandName, subCommandArgs...)
	subCommand.Stdout = os.Stdout

	if err := subCommand.Run(); err != nil {
		return err
	}

	return nil
}
