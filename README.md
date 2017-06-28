# roly

Set AWS environment variables based on your profile.
The goal is to make it easy to work with AWS profiles using assume role.
By working this way you don't need to have credentials in your accounts
if you can assume a role to it from a different account.

## Installation

Using homebrew

```
brew tap diasjorge/tap
brew install roly
```

## Configuration

In your ~/.aws/credentials you can set the profiles like this:

```
[identity-profile]
aws_access_key_id=AWS_ACCESS_KEY_ID
aws_secret_access_key=AWS_SECRET_ACCESS_KEY

[target-profile]
source_profile = identity-profile
mfa_serial = arn:aws:iam::ACCOUNTID:mfa/MFA_DEVICE # Optional only if you need MFA
role_arn = arn:aws:iam::ACCOUNTID:role/SomeRoleName
```

## Usage

roly export target-profile

role exec target-profile env

## Known Limitations

If you use MFA, you the process will stop to ask your MFA on the
stdin, because of that if you pipe the output to another command you
won't see the request message for it and the program will appear to
freeze.
