package aws

import (
    "encoding/json"
    "github.com/google/uuid"
    "strconv"
    "os"
    "emrPricingAPI/models"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/pricing"
)

func commonFetchAwsPriceList(region string, filter *pricing.GetProductsInput) (*pricing.GetProductsOutput, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, err
	}

	// Create a Pricing service client using the custom session
	pricingSvc := pricing.New(sess)

	// Get the AWS price list
	result, err := pricingSvc.GetProducts(filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func commonFetchPricingData(region string, filter *pricing.GetProductsInput) ([]models.Price, error) {
	result, err := commonFetchAwsPriceList(region, filter)
	if err != nil {
		return nil, err
	}

	// Extract the price per unit cost for each instance type
	prices := extractPricingInformation(result)
	return prices, nil
}


//mostly used for debugging
func FetchPricingDataJson(region string, serviceCode string) (*pricing.GetProductsOutput, error) {
	filter := getInputNoFilter(serviceCode)
	result, err := commonFetchAwsPriceList(region, filter)
	if err != nil {
		return nil, err
	}

	saveResponseToJson(result)
	return result, nil
}

func FetchPricingDataJsonFilter(region string, serviceCode string, location string, instanceType string) (*pricing.GetProductsOutput, error) {
	filter := defineFilter(serviceCode, location, instanceType)
	result, err := commonFetchAwsPriceList(region, filter)
	if err != nil {
		return nil, err
	}

	saveResponseToJson(result)
	return result, nil
}

func FetchPricingData(region string, serviceCode string) ([]models.Price, error) {
	filter := getInputNoFilter(serviceCode)
    return commonFetchPricingData(region, filter)
}

func FetchPricingDataFilter(region string, serviceCode string, location string, instanceType string) ([]models.Price, error) {
	filter := defineFilter(serviceCode, location, instanceType)
    return commonFetchPricingData(region, filter)
}

func extractPricingInformation(result *pricing.GetProductsOutput) []models.Price {
    var prices []models.Price

    for _, priceListItem := range result.PriceList {
        price := extractSinglePrice(priceListItem)
        prices = append(prices, price)
    }

    return prices
}

func extractSinglePrice(priceListItem map[string]interface{}) models.Price {
    // Extract product attributes
    productAttributes, _ := priceListItem["product"].(map[string]interface{})["attributes"].(map[string]interface{})
    instanceType, _ := extractInstanceType(productAttributes)

    serviceCode := productAttributes["servicecode"].(string)

    // Initialize price variables
    var p models.Price

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

                p = models.Price{
                    ID:              uuid.New().String(), // Set the ID as needed
                    ServiceType:     serviceCode,
                    InstanceType:    instanceType,
                    Market:          reservationType,
                    Unit:            "USD",
                    PricePerUnit:    pricePerUnitFloat,
                    PriceDescription: priceDescription,
                    UpdatedAt:       "", // Set the update timestamp as needed
                }
            }
        }
    }

    return p
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

//useful for debugging
func saveResponseToJson(result *pricing.GetProductsOutput) error {
    jsonData, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        return err
    }

    // Write the JSON data to a file
    outputFile := "pricingResponse.json"
    err = os.WriteFile(outputFile, jsonData, 0644)
    if err != nil {
        return err
    }

    return nil
}

func getInputNoFilter(serviceCode string) *pricing.GetProductsInput {
    input := &pricing.GetProductsInput{
        		ServiceCode: aws.String(serviceCode),
    }
    return input
}

func defineFilter(serviceCode string, location string, instanceType string) *pricing.GetProductsInput {
    input := &pricing.GetProductsInput{
    		ServiceCode: aws.String(serviceCode),
    		Filters: []*pricing.Filter{
    			{
    				Type:  aws.String("TERM_MATCH"),
    				Field: aws.String("location"),
    				Value: aws.String(location),
    			},
    			{
    				Type:  aws.String("TERM_MATCH"),
    				Field: aws.String("instanceType"),
    				Value: aws.String(instanceType), // Add the instance type filter here
    			},
    		},
    }
    return input
}