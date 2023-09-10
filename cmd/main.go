package main

import (
	"log"

	"github.com/cloud-craftman-go/aws-iam-reports/internal/csvutils"
	"github.com/cloud-craftman-go/aws-iam-reports/internal/setup"
)

func main() {
	app, err := setup.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = csvutils.AWSReport(app.IamClient, *app.Profile)
	if err != nil {
		log.Fatal(err)
	}
}
