package price

import(
 "testing"
 "github.com/stretchr/testify/assert"
 "github.com/claredavies/emrPricingAPI/constants"
 "strings"
 "github.com/claredavies/emrPricingAPI/models"
 )

func TestGetPriceHappyPath(t *testing.T) {
    serviceCode := "ElasticMapReduce"
    instanceType := "m5.2xlarge"
    p, err := GetPrice(serviceCode, instanceType)

    assert.NoError(t, err)
    assert.Equal(t, p.InstanceType, instanceType)
    assert.Equal(t, p.ServiceType, serviceCode)
}

func TestGetPriceBlankServiceCode(t *testing.T) {
    serviceCode := ""
    instanceType := "m5.2xlarge"
    emptyPrice := models.Price{}
    p, err := GetPrice(serviceCode, instanceType)

    assert.Error(t, err)
    assert.True(t, strings.Contains(err.Error(), constants.ErrMsgQueryGetPrice))
    assert.Equal(t, p, emptyPrice)
}

func TestGetPriceBlankInstanceType(t *testing.T) {
    serviceCode := "ElasticMapReduce"
    instanceType := ""
    emptyPrice := models.Price{}
    p, err := GetPrice(serviceCode, instanceType)

    assert.Error(t, err)
    assert.True(t, strings.Contains(err.Error(), constants.ErrMsgQueryGetPrice))
    assert.Equal(t, p, emptyPrice)
}

func TestGetPriceInvalidInstanceType(t *testing.T) {
    serviceCode := "ElasticMapReduce"
    instanceType := "m30.20xlarge"
    emptyPrice := models.Price{}
    p, err := GetPrice(serviceCode, instanceType)

    assert.Error(t, err)
    assert.True(t, strings.Contains(err.Error(), constants.ErrMsgNoMatchingResultsGetPrice))
    assert.Equal(t, p, emptyPrice)
}