package price

import (
	"net/http"
	"testing"
	"fmt"
    "emrPricingAPI/constants"
// 	"encoding/json"
// 	"emrPricingAPI/models"

	"github.com/stretchr/testify/assert"
)

func TestGetPrices_HappyPath(t *testing.T) {
    fmt.Println("--------TestGetPrices_HappyPath-----------")
	_, response, c := getMockRequestResponseContext(http.MethodGet, "/prices")

	if assert.NoError(t, getPrices(c)) {
		assert.Equal(t, http.StatusOK, response.Code)
		responsePrices := getMockResponsePrices(t, response)
		assert.Equal(t, prices, responsePrices)
	}
}

func TestGetPriceEMR_HappyPath(t *testing.T) {
    fmt.Println("--------TestGetPriceEMR_HappyPath-----------")
    queryParams := map[string]string{
        "region": "us-east-1",
        "serviceCode": "ElasticMapReduce",
        "regionCode": "eu-west-1",
        "instanceType": "c6g.12xlarge",
    }
    _, response, c := getMockRequestResponseContextWithQueries(http.MethodGet, "/price", queryParams)

    if assert.NoError(t, getPrice(c)) {
        assert.Equal(t, http.StatusOK, response.Code)
        responsePrice := getMockResponsePrice(t, response)
        fmt.Println(responsePrice)
        assert.Equal(t, queryParams["serviceCode"], responsePrice.ServiceType)
        assert.Equal(t, queryParams["instanceType"], responsePrice.InstanceType)
    }
}

func TestGetPriceEC2_HappyPath(t *testing.T) {
    fmt.Println("--------TestGetPriceEC2_HappyPath-----------")
    queryParams := map[string]string{
        "region": "us-east-1",
        "serviceCode": "AmazonEC2",
        "regionCode": "eu-west-1",
        "instanceType": "c6g.12xlarge",
    }
    _, response, c := getMockRequestResponseContextWithQueries(http.MethodGet, "/price", queryParams)

    if assert.NoError(t, getPrice(c)) {
        assert.Equal(t, http.StatusOK, response.Code)
        responsePrice := getMockResponsePrice(t, response)
        fmt.Println(responsePrice)
        assert.Equal(t, queryParams["serviceCode"], responsePrice.ServiceType)
        assert.Equal(t, queryParams["instanceType"], responsePrice.InstanceType)
    }
}

func TestGetPrice_NoQueryParameter(t *testing.T) {
    fmt.Println("--------TestGetBookById_NoID-----------")
    _, response, c := getMockRequestResponseContext(http.MethodGet, "/price")

    if assert.Error(t, getPrice(c)) {
        assert.Equal(t, http.StatusBadRequest, response.Code)
        responseError := getMockResponseError(t, response)
        assert.Equal(t, constants.ErrMsgQueryGetPrice, responseError["message"])
    }
}

func TestGetPrice_OnlySomeRequiredParams(t *testing.T) {
    fmt.Println("--------TestGetPrice_OnlySomeParams-----------")
    queryParams := map[string]string{
        "region": "us-east-1",
        "serviceCode": "ElasticMapReduce",
        "regionCode": "eu-west-1",
    }
    _, response, c := getMockRequestResponseContextWithQueries(http.MethodGet, "/price", queryParams)

    if assert.Error(t, getPrice(c)) {
            assert.Equal(t, http.StatusBadRequest, response.Code)
            responseError := getMockResponseError(t, response)
            assert.Equal(t, constants.ErrMsgQueryGetPrice, responseError["message"])
    }
}

//no matching results to query params
func TestGetPrice_NoResultsForQueryParamInvalidRegion(t *testing.T) {
    fmt.Println("--------TestGetPrice_NoResultsForQueryParamInvalidRegion-----------")
    queryParams := map[string]string{
        "region": "us-east",
        "serviceCode": "ElasticMapReduce",
        "regionCode": "eu-west-1",
        "instanceType": "c6g.12xlarge",
    }
    _, response, c := getMockRequestResponseContextWithQueries(http.MethodGet, "/price", queryParams)

    if assert.Error(t, getPrice(c)) {
            assert.Equal(t, http.StatusBadRequest, response.Code)
            responseError := getMockResponseError(t, response)
            assert.Equal(t, constants.ErrMsgNoMatchingResultsGetPrice, responseError["message"])
    }
}

func TestGetPrice_NoResultsForQueryParamInvalidInstanceType(t *testing.T) {
    fmt.Println("--------TestGetPrice_NoResultsForQueryParamInvalidInstanceType-----------")
    queryParams := map[string]string{
        "region": "us-east",
        "serviceCode": "ElasticMapReduce",
        "regionCode": "eu-west-1",
        "instanceType": "c9ghdhdxlarge",
    }
    _, response, c := getMockRequestResponseContextWithQueries(http.MethodGet, "/price", queryParams)

    if assert.Error(t, getPrice(c)) {
            assert.Equal(t, http.StatusBadRequest, response.Code)
            responseError := getMockResponseError(t, response)
            assert.Equal(t, constants.ErrMsgNoMatchingResultsGetPrice, responseError["message"])
    }
}

// func TestCreatePrice_HappyPath(t *testing.T) {
//     fmt.Println("--------TestCreatePrice_HappyPath-----------")
//
//     // Create a new book object to be sent in the POST request
//     addPrice := models.AddPrice{
//         Region:       "us-east-1",
//         ServiceCode: "ElasticMapReduce",
//         "regionCode": "eu-west-1",
//         InstanceType: "c6g.12xlarge",
//     }
//
//     // Convert the new book object to JSON
//     requestBody, err := json.Marshal(addPrice)
//     if err != nil {
//         t.Fatal(err)
//     }
//
//     // Perform a POST request to create the new book
//     _, response, c := getMockRequestResponseContextWithRequestBody(http.MethodPost, "/prices", requestBody)
//
//     if assert.NoError(t, createPrice(c)) {
//         // Check if the response status code is as expected (e.g., http.StatusCreated)
//         assert.Equal(t, http.StatusCreated, response.Code)
//
//         createdPrice := getMockResponsePrice(t, response)
//         assert.Equal(t, addPrice.ServiceCode, createdPrice.ServiceType)
//         assert.Equal(t, addPrice.InstanceType, createdPrice.InstanceType)
//
//         retrievedPrice, _ := getPriceByID(createdPrice.ID)
//         assert.Equal(t, retrievedPrice, createdPrice)
//     }
// }
//
// func TestCreatePrice_InvalidRegion(t *testing.T) {
//     fmt.Println("--------TestCreatePrice_InvalidRegion-----------")
//
//     // Create a new book object to be sent in the POST request
//     addPrice := models.AddPrice{
//         Region:       "us-eas",
//         ServiceCode: "ElasticMapReduce",
//         Location:      "US West (Oregon)",
//         InstanceType: "c6g.12xlarge",
//     }
//
//     // Convert the new book object to JSON
//     requestBody, err := json.Marshal(addPrice)
//     if err != nil {
//         t.Fatal(err)
//     }
//
//     // Perform a POST request to create the new book
//     _, response, c := getMockRequestResponseContextWithRequestBody(http.MethodPost, "/prices", requestBody)
//
//     if assert.Error(t, createPrice(c)) {
//         assert.Equal(t, http.StatusBadRequest, response.Code)
//         responseError := getMockResponseError(t, response)
//         assert.Equal(t, constants.ErrMsgNoMatchingResultsGetPrice, responseError["message"])
//     }
// }
//
// func TestCreatePrice_InstanceType(t *testing.T) {
//     fmt.Println("--------TestCreatePrice_InstanceType-----------")
//
//     // Create a new book object to be sent in the POST request
//     addPrice := models.AddPrice{
//         Region:       "us-east-1",
//         ServiceCode: "ElasticMapReduce",
//         Location:      "US West (Oregon)",
//         InstanceType: "c9sksks",
//     }
//
//     // Convert the new book object to JSON
//     requestBody, err := json.Marshal(addPrice)
//     if err != nil {
//         t.Fatal(err)
//     }
//
//     // Perform a POST request to create the new book
//     _, response, c := getMockRequestResponseContextWithRequestBody(http.MethodPost, "/prices", requestBody)
//
//     if assert.NoError(t, createPrice(c)) {
//         assert.Equal(t, http.StatusBadRequest, response.Code)
//         responseError := getMockResponseError(t, response)
//         assert.Equal(t, constants.ErrMsgNoMatchingResultsGetPrice, responseError["message"])
//     }
// }