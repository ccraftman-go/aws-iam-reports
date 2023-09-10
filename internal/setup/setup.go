package setup

import "github.com/aws/aws-sdk-go-v2/service/iam"

type App struct {
	IamClient *iam.Client
	Profile   *string
	Region    *string
}
