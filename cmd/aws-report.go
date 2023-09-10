package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloud-craftman-go/aws-iam-reports/internal/setup"
)

type AWSUser struct {
	User               string `csv:"user"`
	CreationTime       string `csv:"user_creation_time"`
	PassLastUsed       string `csv:"password_last_used"`
	AccessKey1LastUsed string `csv:"access_key_1_last_used_date"`
	AccessKey2LastUsed string `csv:"access_key_2_last_used_date"`
}

func generateAWSReport(app *setup.App) error {
	today := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("aws-report-%s-%s.csv", *app.Profile, today)

	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("unable to create the file: %v", err)
	}
	report, err := app.IamClient.GetCredentialReport(context.TODO(), &iam.GetCredentialReportInput{})
	if err != nil {
		log.Fatalf("unable to retrieve the credentials report, %v", err)
	}

	r := csv.NewReader(bytes.NewReader(report.Content))

	_, err = r.Read()
	if err != nil {
		log.Fatalf("unable to read the header of the CSV: %v", err)
		return err
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"username", "creation_time", "password_last_used", "access_key_1_last_used", "access_key_2_last_used"})

	for {
		row, err := r.Read()
		if err != nil {
			break
		}

		user := AWSUser{
			User:               row[0],
			CreationTime:       row[2],
			PassLastUsed:       row[4],
			AccessKey1LastUsed: row[10],
			AccessKey2LastUsed: row[14],
		}

		writer.Write([]string{
			user.User,
			user.CreationTime,
			user.PassLastUsed,
			user.AccessKey1LastUsed,
			user.AccessKey2LastUsed})
	}

	fmt.Printf("CSV file %s has been written.\n", filename)

	return nil
}
