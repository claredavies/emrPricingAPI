package constants

import "errors"

var (
    ErrQueryParameterMissing = errors.New(ErrMsgQueryGetPrice)
    ErrNoMatchingResults     = errors.New(ErrMsgNoMatchingResultsGetPrice)
)

const (
    ErrMsgPriceNotFound = "Price not found with that ID"
    ErrMsgParamIDRequired    = "ID parameter is required."
    ErrMsgQueryIDRequired    = "Query ID parameter is required."
    ErrInvalidJSON = "Invalid JSON"
    ErrMsgTooManyResultsReturned = "There should only be 1 price matching the query"
    ErrMsgQueryGetPrice    = "Need query parameters: region, serviceCode, regionCode & instanceType"
    ErrMsgNoMatchingResultsGetPrice    = "No matching results for region, serviceCode, regionCode & instanceType"
    ErrMsgBodyAddPrice    = "Need body parameters: region, serviceCode, location & instanceType"
    ErrMsgQueryGetPrices    = "Need query parameters: region & serviceCode"
    ErrMsgAddPrice    = "Price matching that region, serviceCode, location & instanceType already exists"
)