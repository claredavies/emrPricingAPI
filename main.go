package main

import (
    "github.com/claredavies/emrPricingAPI/apis/price"
    "fmt"
)

func main() {
    p, err := price.GetPrice("ElasticMapReduce", "m5.2xlarge")
    prices := price.GetPrices()
    fmt.Println(p)
    fmt.Println(err)

    fmt.Println(prices)
}