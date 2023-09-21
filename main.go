package main

import (
    "fmt"
	"log"
    "emrPricingAPI/pkg/thirdparty/aws"
)

func main() {
	prices, err := aws.FetchPricingData("us-east-1", "ElasticMapReduce")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(prices)
// 	region string, serviceCode string, location string, instanceType string
	prices, err = aws.FetchPricingDataFilter("us-east-1", "ElasticMapReduce", "US West (Oregon)", "C6g.12xlarge")
    	if err != nil {
    		log.Fatalf("Error: %v", err)
    	}
    fmt.Println(prices)
}