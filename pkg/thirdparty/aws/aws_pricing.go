package aws

import (
    "encoding/json"
    "fmt"
    "os"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/pricing"
)

func FetchAndSavePricingData() error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // Replace with your desired region
	})
	if err != nil {
		return err
	}

	// Create a Pricing service client using the custom session
	pricingSvc := pricing.New(sess)

	// Define the filter to get EC2 On-Demand prices
	input := defineFilter()

	// Get the AWS price list
	result, err := pricingSvc.GetProducts(input)
	if err != nil {
		return err
	}

	// Save the response as JSON
	saveResponseToJson(result)

	// Extract the price per unit cost for each instance type
	extractPricingInformation(result)

	return nil
}

func extractPricingInformation(result *pricing.GetProductsOutput) {
	for _, priceListItem := range result.PriceList {
		// Iterate over both "OnDemand" and "Reserved" terms
		for _, reservationType := range []string{"OnDemand", "Reserved"} {
			terms := priceListItem["terms"].(map[string]interface{})[reservationType]
			if terms == nil {
				continue // Skip if the reservation type is not present for this product
			}

			// Iterate over the terms and extract the pricing information
			for _, product := range terms.(map[string]interface{}) {
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

				fmt.Printf("Reservation Type: %s, Instance Type: %s, Price Per Unit: %s\n", reservationType, instanceType, pricePerUnit)
			}
		}
	}
}

func saveResponseToJson(result *pricing.GetProductsOutput) error {
    jsonData, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        return err
    }

    // Write the JSON data to a file
    outputFile := "pricing_response.json"
    err = os.WriteFile(outputFile, jsonData, 0644)
    if err != nil {
        return err
    }

    return nil
}

func defineFilter() *pricing.GetProductsInput {
    input := &pricing.GetProductsInput{
    		ServiceCode: aws.String("ElasticMapReduce"),
    		Filters: []*pricing.Filter{
    			{
    				Type:  aws.String("TERM_MATCH"),
    				Field: aws.String("location"),
    				Value: aws.String("US West (Oregon)"),
    			},
    			{
    				Type:  aws.String("TERM_MATCH"),
    				Field: aws.String("instanceType"),
    				Value: aws.String("C6g.12xlarge"), // Add the instance type filter here
    			},
    		},
    }
    return input
}