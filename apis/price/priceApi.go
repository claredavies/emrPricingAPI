package price

import (
	"errors"
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

func addPrice(c echo.Context) error {
    var newAddPrice models.AddPrice
    if err := c.Bind(&newAddPrice); err != nil {
        return err
    }

    if errValidAddPrice := validatePrice(newAddPrice); errValidAddPrice != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"message": constants.ErrMsgBodyAddPrice})
    }

    onePrice, _ := aws.FetchPricingDataFilter(newAddPrice.Region, newAddPrice.ServiceCode, newAddPrice.Location, newAddPrice.InstanceType)

    //need to handle if it's empty
    if HasMatchingPrice(onePrice[0], prices) {
        return c.JSON(http.StatusBadRequest, echo.Map{"message": constants.ErrMsgAddPrice})
    } else {
        // Check for duplicates before appending
        if !HasMatchingPrice(onePrice[0], prices) {
            // If not a duplicate, append onePrice[0] to prices
            prices = append(prices, onePrice[0])
        }
    }

    err := c.JSON(http.StatusCreated, onePrice)
    if err != nil {
        return err
    }

    return nil
}

func validatePrice(addPrice models.AddPrice) error {
      // Check if the Title is empty
        if addPrice.Region == "" {
            return errors.New("Region cannot be empty")
        }

        // Check if the Author is empty
        if addPrice.ServiceCode == "" {
            return errors.New("ServiceCode cannot be empty")
        }

        if addPrice.Location == "" {
                return errors.New("Location cannot be empty")
            }

        if addPrice.InstanceType == "" {
                return errors.New("InstanceType cannot be empty")
            }

        return nil
}

func HasMatchingPrice(priceToCheck models.Price, prices []models.Price) bool {
    for _, price := range prices {
        if priceToCheck.ServiceType == price.ServiceType &&
            priceToCheck.InstanceType == price.InstanceType &&
            priceToCheck.Market == price.Market {
            return true
        }
    }
    return false
}

func getOnePriceViaQueryParams(c echo.Context) ([]models.Price, error) {
    region := c.QueryParam("region")
    serviceCode := c.QueryParam("serviceCode")
    location := c.QueryParam("location")
    instanceType := c.QueryParam("instanceType")

    if region == "" || serviceCode == "" || location == "" || instanceType == "" {
        return nil, echo.NewHTTPError(http.StatusBadRequest, constants.ErrMsgQueryGetPrice)
    }

    onePrice, err := aws.FetchPricingDataFilter(region, serviceCode, location, instanceType)
    return onePrice, err
}

func getPrice(c echo.Context) error {
    onePrice, err := getOnePriceViaQueryParams(c)

    if err != nil {
        c.JSON(http.StatusBadRequest, echo.Map{"message": constants.ErrMsgQueryGetPrice})
        return err
    }

    return c.JSON(http.StatusOK, onePrice)
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
    e.POST("/addPrice", addPrice)
    e.GET("/unstructuredJson", fetchJsonUnstructured)
    e.GET("/unstructuredJsonFilter", fetchJsonUnstructuredFilter)
}