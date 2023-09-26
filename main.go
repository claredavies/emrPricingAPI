package main

import (
	"github.com/labstack/echo/v4"
    "emrPricingAPI/apis/price"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
    e := echo.New()
    e.Use(middleware.Recover())

    price.SetupRoutes(e)
//     price.LoadPrices("us-east-1", "ElasticMapReduce")
    e.Start("localhost:8080")
}