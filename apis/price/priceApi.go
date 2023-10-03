package price

import (
	"errors"
    "strings"
	"github.com/labstack/echo/v4"
	"net/http"
	"emrPricingAPI/models"
	"emrPricingAPI/pkg/thirdparty/aws"
	"emrPricingAPI/constants"
)

var prices = []models.Price{
}

func getPriceByID(id string) (models.Price, error) {
    for _, p := range prices {
        if p.ID == id {
            return p, nil
        }
    }
    return models.Price{}, errors.New(constants.ErrMsgPriceNotFound)
}

func LoadPrices(region string, service string) {
    newPrices, _ := aws.FetchPricingData(region, service)
    prices = append(prices, newPrices...)
}

func hasExistingInstanceTypeServiceType(serviceCode string,  instanceType string, prices []models.Price) bool {
    for _, price := range prices {
        if serviceCode == price.ServiceType &&
            strings.ToLower(instanceType) == strings.ToLower(price.InstanceType) {
            return true
        }
    }
    return false
}

func getPriceInstanceTypeServiceType(serviceCode string, instanceType string, prices []models.Price) *models.Price {
    for _, price := range prices {
        if serviceCode == price.ServiceType &&
            strings.ToLower(instanceType) == strings.ToLower(price.InstanceType) {
            return &price
        }
    }
    return nil
}

func getOnePriceViaQueryParams(c echo.Context) (models.Price, error) {
    serviceCode := c.QueryParam("serviceCode")
    instanceType := c.QueryParam("instanceType")

    if serviceCode == "" || instanceType == "" {
        return models.Price{}, constants.ErrQueryParameterMissing
    }

    onePrice, errRequestError := aws.FetchPricingDataFilter(constants.Region, serviceCode, constants.RegionCode, instanceType)
    if errRequestError != nil {
        return models.Price{}, constants.ErrNoMatchingResults
    }

    if len(onePrice) == 0 {
            return models.Price{}, constants.ErrNoMatchingResults
        } else if len(onePrice) > 1 {
            return models.Price{}, constants.ErrTooManyResultsReturned
        }

    return onePrice[0], errRequestError
}

func fetchJsonUnstructuredFilter(c echo.Context) error {
    serviceCode:=c.QueryParam("serviceCode")
    instanceType:=c.QueryParam("instanceType")

    if serviceCode == "" || instanceType == "" {
        return c.JSON(http.StatusBadRequest, echo.Map{"message": constants.ErrMsgQueryGetPrice})
    }

    jsonResult, _ := aws.FetchPricingDataJsonFilter(constants.Region, serviceCode, constants.RegionCode, instanceType)
    err := c.JSON(http.StatusOK, jsonResult)
    if err != nil {
        return err
    }

    return nil
}

func getPrice(c echo.Context) error {
    serviceCode := c.QueryParam("serviceCode")
    instanceType := c.QueryParam("instanceType")

    if serviceCode == "" || instanceType == "" {
        return c.JSON(http.StatusBadRequest, echo.Map{"message": constants.ErrMsgQueryGetPrice})
    }

    alreadyExists := hasExistingInstanceTypeServiceType(serviceCode,  instanceType, prices)

    if alreadyExists == false {
        onePrice, err := getOnePriceViaQueryParams(c)
        if err == nil {
           prices = append(prices, onePrice)
           return c.JSON(http.StatusOK, onePrice)
        } else {
            if err == constants.ErrTooManyResultsReturned {
               return c.JSON(http.StatusBadRequest, echo.Map{"message": constants.ErrMsgTooManyResultsReturned})
            } else if err == constants.ErrQueryParameterMissing {
               return c.JSON(http.StatusBadRequest, echo.Map{"message": constants.ErrMsgQueryGetPrice})
            } else {
               return c.JSON(http.StatusBadRequest, echo.Map{"message": constants.ErrMsgNoMatchingResultsGetPrice})
            }
        }
    } else {
        price := getPriceInstanceTypeServiceType(serviceCode,  instanceType, prices)
        return c.JSON(http.StatusOK, price)
    }
}

func getPrices(c echo.Context) error {
	err := c.JSON(http.StatusOK, prices)
    if err != nil {
        return err
    }

    return nil
}

func fetchJsonUnstructured(c echo.Context) error {
    serviceCode :=c.QueryParam("serviceCode")

    if serviceCode == "" {
        return c.JSON(http.StatusBadRequest, echo.Map{"message": constants.ErrMsgQueryGetPrices})
    }

    jsonResult, _ := aws.FetchPricingDataJson(constants.Region, serviceCode)
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