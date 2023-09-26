package price

import (
	"net/http"
	"testing"
	"fmt"
    "emrPricingAPI/constants"
// 	"encoding/json"

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
        "location": "US West (Oregon)",
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
        "location": "US West (Oregon)",
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
        "location": "US West (Oregon)",
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
        "location": "US West (Oregon)",
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
        "location": "US West (Oregon)",
        "instanceType": "c9ghdhdxlarge",
    }
    _, response, c := getMockRequestResponseContextWithQueries(http.MethodGet, "/price", queryParams)

    if assert.Error(t, getPrice(c)) {
            assert.Equal(t, http.StatusBadRequest, response.Code)
            responseError := getMockResponseError(t, response)
            assert.Equal(t, constants.ErrMsgNoMatchingResultsGetPrice, responseError["message"])
    }
}

//multiple results to query params