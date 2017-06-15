package credentials

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
)

var keys = []string{
	"AWS_ACCESS_KEY_ID",
	"AWS_SECRET_ACCESS_KEY",
	"AWS_SESSION_TOKEN",
}

func cleanEnv() error {
	for _, key := range keys {
		if err := os.Setenv(key, ""); err != nil {
			return err
		}
	}
	return nil
}

// Get returns Profile Credentials
func Get(profile string) (credentials.Value, error) {
	if err := cleanEnv(); err != nil {
		return *&credentials.Value{}, err
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			CredentialsChainVerboseErrors: aws.Bool(true),
		},
		Profile:                 profile,
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
		SharedConfigState:       session.SharedConfigEnable,
	}))

	return sess.Config.Credentials.Get()
}
