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

	 // Create a new AWS session with custom configuration
        sess, err := session.NewSession(&aws.Config{
            Region:      aws.String(awsRegion),
            Credentials: credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, awsSessionToken),
        })
        if err != nil {
            fmt.Println("Error creating AWS session:", err)
            return
        }

        // Create a Pricing service client using the custom session
        pricingSvc := pricing.New(sess)

        // Define the filter to get EC2 On-Demand prices
        input := &pricing.GetProductsInput{
            ServiceCode: aws.String("AmazonEC2"),
            Filters: []*pricing.Filter{
                {
                    Type:  aws.String("TERM_MATCH"),
                    Field: aws.String("productFamily"),
                    Value: aws.String("Compute Instance"),
                },
                {
                    Type:  aws.String("TERM_MATCH"),
                    Field: aws.String("tenancy"),
                    Value: aws.String("Shared"),
                },
                {
                    Type:  aws.String("TERM_MATCH"),
                    Field: aws.String("location"),
                    Value: aws.String("US East (N. Virginia)"),
                },
                {
                    Type:  aws.String("TERM_MATCH"),
                    Field: aws.String("operatingSystem"),
                    Value: aws.String("Linux"),
                },
                {
                    Type:  aws.String("TERM_MATCH"),
                    Field: aws.String("instanceType"),
                    Value: aws.String("t2.micro"), // Add the instance type filter here
                },
                // You can add more filters as needed
            },
        }

        // Get the AWS price list
        result, err := pricingSvc.GetProducts(input)
        if err != nil {
            fmt.Println("Error fetching AWS Pricing API data:", err)
            return
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

        // Extract and print the price per unit cost for each instance type
        for _, priceListItem := range result.PriceList {
//             instanceType := priceListItem["product"].(map[string]interface{})["attributes"].(map[string]interface{})["instanceType"].(string)
//
//             // Extract price per unit (modify the path based on the JSON structure)
//             pricePerUnit := priceListItem["terms"].(map[string]interface{})["OnDemand"].(map[string]interface{})["YOUR_PRICE_DIMENSION_KEY"].(map[string]interface{})["pricePerUnit"].(map[string]interface{})["USD"].(string)

            fmt.Printf("Instance Type: %s, Price Per Unit: %s\n", priceListItem["terms"])
        }
    }