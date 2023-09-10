package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloud-craftman-go/aws-iam-reports/internal/setup"
)

func setupApp() (*setup.App, error) {
	profile := flag.String("profile", "", "the aws profile in the .aws/credentials file")
	region := flag.String("region", "us-east-1", "the aws region where the commands will run")

	flag.Parse()

	if *profile == "" {
		fmt.Println("Missing required flags.")
		os.Exit(1)
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(*region),
		config.WithSharedConfigProfile(*profile))

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	iamClient := iam.NewFromConfig(cfg)

	app := &setup.App{
		IamClient: iamClient,
		Profile:   profile,
		Region:    region,
	}

	return app, err
}
