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
		responsePrices := getMockResponseBooks(t, response)
		assert.Equal(t, prices, responsePrices)
	}
}

func TestGetPrice_HappyPath(t *testing.T) {
    fmt.Println("--------TestGetPrice_HappyPath-----------")
    queryParams := map[string]string{
        "region": "us-east-1",
        "serviceCode": "ElasticMapReduce",
        "location": "US West (Oregon)",
        "instanceType": "c6g.12xlarge",
    }
    _, response, c := getMockRequestResponseContextWithQueries(http.MethodGet, "/price", queryParams)

    if assert.NoError(t, getPrice(c)) {
        assert.Equal(t, http.StatusOK, response.Code)
        responsePrice := getMockResponseBooks(t, response)
        fmt.Println(responsePrice)
        assert.Equal(t, queryParams["serviceCode"], responsePrice[0].ServiceType)
        assert.Equal(t, queryParams["instanceType"], responsePrice[0].InstanceType)
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

//multiple results to query params

// func TestReturnBook_NoQueryParameter(t *testing.T) {
//     fmt.Println("--------TestReturnBook_NoQueryParameter-----------")
//     _, response, c := getMockRequestResponseContext(http.MethodPatch, "/return")
//
//     if assert.NoError(t, returnBook(c)) {
//            assert.Equal(t, http.StatusBadRequest, response.Code)
//            responseError := getMockResponseError(t, response)
//            assert.Equal(t, constants.ErrMsgQueryIDRequired, responseError["message"])
//     }
// }
//
// func TestReturnBook_InvalidBookId(t *testing.T) {
//     fmt.Println("--------TestReturnBook_InvalidBookId-----------")
//     _, response, c := getMockRequestResponseContextWithQuery(http.MethodPatch, "/return", "id", "ddd")
//
//     if assert.NoError(t, returnBook(c)) {
//            assert.Equal(t, http.StatusNotFound, response.Code)
//            responseError := getMockResponseError(t, response)
//            assert.Equal(t, constants.ErrMsgBookNotFound, responseError["message"])
//     }
// }