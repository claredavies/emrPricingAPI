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
            ServiceCode: aws.String("ElasticMapReduce"),
            Filters: []*pricing.Filter{
                {
                    Type:  aws.String("TERM_MATCH"),
                    Field: aws.String("location"),
                    Value: aws.String("US East (N. Virginia)"),
                },
                {
                    Type:  aws.String("TERM_MATCH"),
                    Field: aws.String("instanceType"),
                    Value: aws.String("C6g.12xlarge"), // Add the instance type filter here
                },
                {
                    Type:  aws.String("TERM_MATCH"),
                    Field: aws.String("softwareType"),
                    Value: aws.String("EMR"), // Add the instance type filter here
                },
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

        // Extract the price per unit cost for each instance type
        // Extract the desired fields from the JSON response
        for _, priceListItem := range result.PriceList {
            terms := priceListItem["terms"].(map[string]interface{})["OnDemand"].(map[string]interface{})
            for _, product := range terms {
                priceDimensions := product.(map[string]interface{})["priceDimensions"].(map[string]interface{})
                firstNestedElement := map[string]interface{}{} // Initialize an empty map for the first nested element

                // Iterate over the price dimensions to access the first nested element
                for _, dimension := range priceDimensions {
                    firstNestedElement = dimension.(map[string]interface{})
                    break // Break after accessing the first nested element
                }

                // Extract the price per unit cost and instance type from the first nested element
                pricePerUnit := firstNestedElement["pricePerUnit"].(map[string]interface{})["USD"].(string)
                instanceType := firstNestedElement["description"].(string)

                fmt.Printf("Instance Type: %s, Price Per Unit: %s\n", instanceType, pricePerUnit)
            }
        }

}