package csvutils

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

type AWSUser struct {
	User               string `csv:"user"`
	CreationTime       string `csv:"user_creation_time"`
	PassLastUsed       string `csv:"password_last_used"`
	AccessKey1LastUsed string `csv:"access_key_1_last_used_date"`
	AccessKey2LastUsed string `csv:"access_key_2_last_used_date"`
}

func ExportToFile(header []string, content []byte, profile string) error {
	today := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("aws-report-%s-%s.csv", profile, today)

	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("unable to create the file: %v", err)
	}

	r := csv.NewReader(bytes.NewReader(content))

	_, err = r.Read()
	if err != nil {
		log.Fatalf("unable to read the header of the CSV: %v", err)
		return err
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write(header)

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
			AccessKey2LastUsed: row[15],
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
