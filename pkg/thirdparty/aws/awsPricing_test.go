package aws

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFetchPricingDataJson(t *testing.T) {
	_, err := FetchPricingDataJson("us-east-1", "ElasticMapReduce")
	assert.NoError(t, err)
}

func TestFetchPricingData(t *testing.T) {
	prices, err := FetchPricingData("us-east-1", "ElasticMapReduce")
	assert.NoError(t, err)
	assert.NotEmpty(t, prices)
}

func TestFetchPricingDataFilter(t *testing.T) {
	prices, err := FetchPricingDataFilter("us-east-1", "ElasticMapReduce", "US West (Oregon)", "C6g.12xlarge")
	assert.NoError(t, err)
	assert.NotEmpty(t, prices)
}