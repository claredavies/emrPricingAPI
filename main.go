package main

import (
    "encoding/json"
	"fmt"
	"log"
	"os"
    "emrPricingAPI/constants"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/pricing"
)

func main() {
	// Replace these values with your own AWS credentials and region
	awsAccessKeyID := constants.AwsAccessKeyID
	awsSecretAccessKey := constants.AwsSecretAccessKey
	awsSessionToken := constants.AwsSessionToken
	awsRegion := "us-east-1" // Replace with your desired AWS region

	// Create AWS credentials
	creds := credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, awsSessionToken)

    // Create a new AWS session with custom configuration
    sess, err := session.NewSession(&aws.Config{
        Region:      aws.String(awsRegion),
        Credentials: creds,
    })
    if err != nil {
        log.Fatalf("Error creating AWS session: %v", err)
    }

    // Create a Pricing service client using the custom session
    pricingSvc := pricing.New(sess)

    // Make a basic request to the Pricing API
    input := &pricing.GetProductsInput{
        ServiceCode:   aws.String("AmazonEC2"),
        FormatVersion: aws.String("aws_v1"),
    }
    result, err := pricingSvc.GetProducts(input)
    if err != nil {
        log.Fatalf("Error fetching AWS Pricing API data: %v", err)
    }

    // Save the response as JSON
    jsonData, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        log.Fatalf("Error encoding JSON: %v", err)
    }

    // Write the JSON data to a file
    outputFile := "pricing_response.json"
    err = os.WriteFile(outputFile, jsonData, 0644)
    if err != nil {
        log.Fatalf("Error writing JSON to file: %v", err)
    }

    fmt.Printf("Pricing data saved to %s\n", outputFile)
}