package price

import (
// 	"errors"
//     "fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"emrPricingAPI/models"
	"emrPricingAPI/pkg/thirdparty/aws"
	"emrPricingAPI/constants"
)

var prices = []models.Price{
}

//input region and service
func LoadPrices() {
    newPrices, _ := aws.FetchPricingData("us-east-1", "ElasticMapReduce")
    prices = append(prices, newPrices...)
}

func getPrice(c echo.Context) error {
    region:=c.QueryParam("region")
    serviceCode:=c.QueryParam("serviceCode")
    location:=c.QueryParam("location")
    instanceType:=c.QueryParam("instanceType")

    if region == ""|| serviceCode == ""|| location == ""|| instanceType == ""{
        return c.JSON(http.StatusBadRequest, echo.Map{"message": constants.ErrMsgQueryGetPrice})
    }

    onePrice, _ := aws.FetchPricingDataFilter(region, serviceCode, location, instanceType)
    err := c.JSON(http.StatusOK, onePrice)
    if err != nil {
        return err
    }

    return nil
}

func fetchJsonUnstructured(c echo.Context) error {
    jsonResult, _ := aws.FetchPricingDataJson("us-east-1", "ElasticMapReduce")
    err := c.JSON(http.StatusOK, jsonResult)
    if err != nil {
        return err
    }

    return nil
}

func fetchJsonUnstructuredFilter(c echo.Context) error {
    jsonResult, _ := aws.FetchPricingDataJsonFilter("us-east-1", "ElasticMapReduce", "US West (Oregon)", "C6g.12xlarge")
    err := c.JSON(http.StatusOK, jsonResult)
    if err != nil {
        return err
    }

    return nil
}

func getPrices(c echo.Context) error {
	err := c.JSON(http.StatusOK, prices)
    if err != nil {
        return err
    }

    return nil
}

func SetupRoutes(e *echo.Echo) {
    e.GET("/prices", getPrices)
    e.GET("/getPrice", getPrice)
    e.GET("/unstructuredJson", fetchJsonUnstructured)
    e.GET("/unstructuredJsonFilter", fetchJsonUnstructuredFilter)
}