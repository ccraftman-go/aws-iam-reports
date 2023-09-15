package main

import (
	"log"

	"github.com/ccraftman-go/aws-iam-reports/internal/reports"
	"github.com/ccraftman-go/aws-iam-reports/internal/setup"
)

func main() {
	app, err := setup.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = reports.AWSReport(app.IamClient, *app.Profile)
	if err != nil {
		log.Fatal(err)
	}
}
