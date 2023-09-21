package main

import (
	"github.com/labstack/echo/v4"
    "emrPricingAPI/apis/price"
)

func main() {
    e := echo.New()
    price.SetupRoutes(e)
    price.LoadPrices()
    e.Start("localhost:8080")
}