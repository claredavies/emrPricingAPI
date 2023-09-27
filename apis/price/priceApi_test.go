package price

import (
	"net/http"
	"testing"
	"fmt"
    "emrPricingAPI/constants"
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
        "serviceCode": "ElasticMapReduce",
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
        "serviceCode": "AmazonEC2",
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
    fmt.Println("--------TestGetPrice_NoQueryParameter-----------")
    _, response, c := getMockRequestResponseContext(http.MethodGet, "/price")

    if assert.NoError(t, getPrice(c)) {
        assert.Equal(t, http.StatusBadRequest, response.Code)
        responseError := getMockResponseError(t, response)
        assert.Equal(t, constants.ErrMsgQueryGetPrice, responseError["message"])
    }
}

func TestGetPrice_OnlySomeRequiredParams(t *testing.T) {
    fmt.Println("--------TestGetPrice_OnlySomeParams-----------")
    queryParams := map[string]string{
        "serviceCode": "ElasticMapReduce",
    }
    _, response, c := getMockRequestResponseContextWithQueries(http.MethodGet, "/price", queryParams)

    if assert.NoError(t, getPrice(c)) {
            assert.Equal(t, http.StatusBadRequest, response.Code)
            responseError := getMockResponseError(t, response)
            assert.Equal(t, constants.ErrMsgQueryGetPrice, responseError["message"])
    }
}


func TestGetPrice_NoResultsForQueryParamInvalidInstanceType(t *testing.T) {
    fmt.Println("--------TestGetPrice_NoResultsForQueryParamInvalidInstanceType-----------")
    queryParams := map[string]string{
        "serviceCode": "ElasticMapReduce",
        "instanceType": "c9ghdhdxlarge",
    }
    _, response, c := getMockRequestResponseContextWithQueries(http.MethodGet, "/price", queryParams)

    if assert.NoError(t, getPrice(c)) {
            assert.Equal(t, http.StatusBadRequest, response.Code)
            responseError := getMockResponseError(t, response)
            assert.Equal(t, constants.ErrMsgNoMatchingResultsGetPrice, responseError["message"])
    }
}