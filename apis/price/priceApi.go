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

func LoadPrices(region string, service string) {
    newPrices, _ := aws.FetchPricingData(region, service)
    prices = append(prices, newPrices...)
}

func getOnePrice(c echo.Context) ([]models.Price, error) {
    region:=c.QueryParam("region")
    serviceCode:=c.QueryParam("serviceCode")
    location:=c.QueryParam("location")
    instanceType:=c.QueryParam("instanceType")

    if region == ""|| serviceCode == ""|| location == ""|| instanceType == ""{
        return nil, c.JSON(http.StatusBadRequest, echo.Map{"message": constants.ErrMsgQueryGetPrice})
    }

    onePrice, err := aws.FetchPricingDataFilter(region, serviceCode, location, instanceType)
    return onePrice, err
}

func getPrice(c echo.Context) error {
    onePrice, _ := getOnePrice(c)
    err := c.JSON(http.StatusOK, onePrice)
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

func fetchJsonUnstructured(c echo.Context) error {
    region :=c.QueryParam("region")
    serviceCode :=c.QueryParam("serviceCode")

    if region == ""|| serviceCode == "" {
            return c.JSON(http.StatusBadRequest, echo.Map{"message": constants.ErrMsgQueryGetPrices})
    }

    jsonResult, _ := aws.FetchPricingDataJson(region, serviceCode)
    err := c.JSON(http.StatusOK, jsonResult)
    if err != nil {
        return err
    }

    return nil
}

func fetchJsonUnstructuredFilter(c echo.Context) error {
    region:=c.QueryParam("region")
    serviceCode:=c.QueryParam("serviceCode")
    location:=c.QueryParam("location")
    instanceType:=c.QueryParam("instanceType")

    if region == ""|| serviceCode == ""|| location == ""|| instanceType == ""{
        return c.JSON(http.StatusBadRequest, echo.Map{"message": constants.ErrMsgQueryGetPrice})
    }

    jsonResult, _ := aws.FetchPricingDataJsonFilter(region, serviceCode, location, instanceType)
    err := c.JSON(http.StatusOK, jsonResult)
    if err != nil {
        return err
    }

    return nil
}

func SetupRoutes(e *echo.Echo) {
    e.GET("/prices", getPrices)
    e.GET("/price", getPrice)
    e.GET("/unstructuredJson", fetchJsonUnstructured)
    e.GET("/unstructuredJsonFilter", fetchJsonUnstructuredFilter)
}