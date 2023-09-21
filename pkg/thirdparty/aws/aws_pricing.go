package aws

import (
    "encoding/json"
    "fmt"
    "strconv"
    "os"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/pricing"
)

type Price struct {
    ID              string  `json:"ID"`
    ServiceType     string  `json:"ServiceType"`
    InstanceType    string  `json:"InstanceType"`
    Market          string  `json:"Market"`
    Unit            string  `json:"Unit"`
    PricePerUnit    float64 `json:"PricePerUnit"`
    PriceDescription string  `json:"PriceDescription"`
    UpdatedAt       string  `json:"UpdatedAt"`
}

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
	prices := extractPricingInformation(result)
	fmt.Println(prices)

	return nil
}

func extractPricingInformation(result *pricing.GetProductsOutput) []Price {
    var prices []Price

    for _, priceListItem := range result.PriceList {
        // Extract product attributes
        productAttributes, _ := priceListItem["product"].(map[string]interface{})["attributes"].(map[string]interface{})
        instanceType, _ := extractInstanceType(productAttributes)

        serviceCode := productAttributes["servicecode"].(string)
//         regionCode := productAttributes["regionCode"].(string)

        // Iterate over both "OnDemand" and "Reserved" terms
        for _, reservationType := range []string{"OnDemand", "Reserved"} {
            terms := priceListItem["terms"].(map[string]interface{})[reservationType]
            if terms == nil {
                continue // Skip if the reservation type is not present for this product
            }

            // Iterate over the terms and extract the pricing information
            for _, product := range terms.(map[string]interface{}) {
                priceDimensions := product.(map[string]interface{})["priceDimensions"].(map[string]interface{})
                for _, dimension := range priceDimensions {
                    pricePerUnit, _ := dimension.(map[string]interface{})["pricePerUnit"].(map[string]interface{})["USD"].(string)
                    pricePerUnitFloat, _ := strconv.ParseFloat(pricePerUnit, 64)
                    priceDescription := dimension.(map[string]interface{})["description"].(string)

                    price := Price{
                        ID:              "", // Set the ID as needed
                        ServiceType:     serviceCode,
                        InstanceType:    instanceType,
                        Market:          reservationType,
                        Unit:            "USD",
                        PricePerUnit:    pricePerUnitFloat,
                        PriceDescription: priceDescription,
                        UpdatedAt:       "", // Set the update timestamp as needed
                    }

                    prices = append(prices, price)
                }
            }
        }
    }

    return prices
}

func extractInstanceType(productAttributes map[string]interface{}) (string, bool) {
    fieldNames := []string{"instanceType", "compute", "usagetype"}

    for _, fieldName := range fieldNames {
        instanceType, ok := productAttributes[fieldName].(string)
        if ok {
            return instanceType, true
        }
    }

    return "", false
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
//     			{
//     				Type:  aws.String("TERM_MATCH"),
//     				Field: aws.String("instanceType"),
//     				Value: aws.String("C6g.12xlarge"), // Add the instance type filter here
//     			},
    		},
    }
    return input
}