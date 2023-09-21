package price

import (
// 	"errors"
//     "fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"emrPricingAPI/models"
	"emrPricingAPI/pkg/thirdparty/aws"
)

var prices = []models.Price{
}

func LoadPrices() {
    newPrices, _ := aws.FetchPricingData("us-east-1", "ElasticMapReduce")
    prices = append(prices, newPrices...)
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
}