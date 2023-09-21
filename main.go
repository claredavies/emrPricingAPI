package main

import (
	"log"
    "emrPricingAPI/pkg/thirdparty/aws"
)

func main() {
	err := aws.FetchAndSavePricingData()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}