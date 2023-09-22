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

func fetchOnePrice(c echo.Context) error {
    onePrice, _ := aws.FetchPricingDataFilter("us-east-1", "ElasticMapReduce", "US West (Oregon)", "C6g.12xlarge")
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

func SetupRoutes(e *echo.Echo) {
    e.GET("/prices", getPrices)
    e.GET("/onePrice", fetchOnePrice)
}