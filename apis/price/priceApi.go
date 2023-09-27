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

func createPrice(c echo.Context) error {
    var newAddPrice models.AddPrice
    // Bind the request body to newAddPrice
    if err := c.Bind(&newAddPrice); err != nil {
        return err
    }

    // Validate the newAddPrice
    if err := validatePrice(newAddPrice); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"message": constants.ErrMsgBodyAddPrice})
    }

    // Fetch pricing data based on the newAddPrice parameters
    onePrice, err := aws.FetchPricingDataFilter(newAddPrice.Region, newAddPrice.ServiceCode, newAddPrice.Location, newAddPrice.InstanceType)

    if err != nil {
        c.JSON(http.StatusBadRequest, echo.Map{"message": constants.ErrMsgNoMatchingResultsGetPrice})
        return err
    }

    if len(onePrice) == 0 {
        return c.JSON(http.StatusBadRequest, echo.Map{"message": constants.ErrMsgNoMatchingResultsGetPrice})
    } else if len(onePrice) == 1 {
        if HasMatchingPrice(onePrice[0], prices) {
            return c.JSON(http.StatusBadRequest, echo.Map{"message": constants.ErrMsgAddPrice})
        }

        // Check for duplicates before appending
        if !HasMatchingPrice(onePrice[0], prices) {
            // If not a duplicate, append onePrice[0] to prices
            prices = append(prices, onePrice[0])
        }
    } else {
        return c.JSON(http.StatusInternalServerError, echo.Map{"message": constants.ErrMsgTooManyResultsReturned})
    }

    // Return a JSON response with the created price
    return c.JSON(http.StatusCreated, onePrice[0])
}

func validatePrice(addPrice models.AddPrice) error {
        if addPrice.ServiceCode == "" {
            return errors.New("ServiceCode cannot be empty")
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
    serviceCode := c.QueryParam("serviceCode")
    instanceType := c.QueryParam("instanceType")

    if serviceCode == "" || instanceType == "" {
        return nil, constants.ErrQueryParameterMissing
    }

    onePrice, errRequestError := aws.FetchPricingDataFilter(constants.Region, serviceCode, constants.RegionCode, instanceType)

    if errRequestError != nil {
        return nil, constants.ErrNoMatchingResults
    }

    return onePrice, errRequestError
}

func getPrice(c echo.Context) error {
    onePrice, err := getOnePriceViaQueryParams(c)

    if err != nil {
        switch err.(type) {
        case *echo.HTTPError:
            return err
        case error:
            if err == constants.ErrQueryParameterMissing {
                c.JSON(http.StatusBadRequest, echo.Map{"message": constants.ErrMsgQueryGetPrice})
                return err
            } else if err == constants.ErrNoMatchingResults {
                c.JSON(http.StatusBadRequest, echo.Map{"message": constants.ErrMsgNoMatchingResultsGetPrice})
                return err
            }
            return err
        }
    }

    if len(onePrice) == 0 {
        return c.JSON(http.StatusBadRequest, echo.Map{"message": constants.ErrMsgNoMatchingResultsGetPrice})
    } else if len(onePrice) > 1 {
        return c.JSON(http.StatusInternalServerError, echo.Map{"message": constants.ErrMsgTooManyResultsReturned})
    }

    return c.JSON(http.StatusOK, onePrice[0])
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

func SetupRoutes(e *echo.Echo) {
    e.GET("/prices", getPrices)
    e.GET("/price", getPrice)
    e.POST("/prices", createPrice)
    e.GET("/unstructuredJson", fetchJsonUnstructured)
    e.GET("/unstructuredJsonFilter", fetchJsonUnstructuredFilter)
}