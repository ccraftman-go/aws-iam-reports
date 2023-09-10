package main

import (
	"log"
)

func main() {
	app, err := setupApp()
	if err != nil {
		log.Fatal(err)
	}

	err = generateAWSReport(app)
	if err != nil {
		log.Fatal(err)
	}
}
