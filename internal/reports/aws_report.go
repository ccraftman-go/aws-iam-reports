package reports

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/ccraftman-go/aws-iam-reports/internal/csvutils"
)

func AWSReport(iamClient *iam.Client, profile string) error {

	_, err := iamClient.GenerateCredentialReport(context.TODO(), &iam.GenerateCredentialReportInput{})
	if err != nil {
		log.Fatalf("unable to generate the credentials report, %v", err)
	}

	report, err := waitForCredentialReport(iamClient, 10*time.Second)
	if err != nil {
		log.Fatalf("unable to retrieve the credentials report: %v", err)
	}

	header := []string{"username", "creation_time", "password_last_used", "access_key_1_last_used", "access_key_2_last_used"}

	csvutils.ExportToFile(header, report.Content, profile)

	return nil
}

func waitForCredentialReport(iamClient *iam.Client, timeout time.Duration) (*iam.GetCredentialReportOutput, error) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	timeoutCh := time.After(timeout)

	for {
		select {
		case <-ticker.C:
			genReport, genErr := iamClient.GenerateCredentialReport(context.TODO(), &iam.GenerateCredentialReportInput{})
			if genErr != nil {
				return nil, genErr
			}
			fmt.Println(genReport.State)
			if genReport.State == "COMPLETE" {
				report, err := iamClient.GetCredentialReport(context.TODO(), &iam.GetCredentialReportInput{})
				if err == nil && report != nil {
					return report, nil
				}
			}
		case <-timeoutCh:
			return nil, fmt.Errorf("timed out waiting for credential report")
		}
	}
}
