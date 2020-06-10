package credentials

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Get returns Profile Credentials
func Get(profile string, quiet bool) (credentials.Value, error) {
	stscreds.DefaultDuration = 3600 * time.Second

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			CredentialsChainVerboseErrors: aws.Bool(true),
		},
		Profile: profile,
		AssumeRoleTokenProvider: func() (string, error) {
			var v string
			if !quiet {
				fmt.Printf("Assume Role MFA token code: ")
			}
			_, err := fmt.Scanln(&v)
			return v, err
		},
		SharedConfigState: session.SharedConfigEnable,
	}))

	return sess.Config.Credentials.Get()
}
