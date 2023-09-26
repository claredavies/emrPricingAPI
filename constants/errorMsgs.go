package constants

import "errors"

var (
    ErrQueryParameterMissing = errors.New(ErrMsgQueryGetPrice)
    ErrNoMatchingResults     = errors.New(ErrMsgNoMatchingResultsGetPrice)
)

const (
    ErrMsgParamIDRequired    = "ID parameter is required."
    ErrMsgQueryIDRequired    = "Query ID parameter is required."
    ErrInvalidJSON = "Invalid JSON"
    ErrMsgTooManyResultsReturned = "There should only be 1 price matching the query"
    ErrMsgQueryGetPrice    = "Need query parameters: region, serviceCode, location & instanceType"
    ErrMsgNoMatchingResultsGetPrice    = "No matching results for region, serviceCode, location & instanceType"
    ErrMsgBodyAddPrice    = "Need body parameters: region, serviceCode, location & instanceType"
    ErrMsgQueryGetPrices    = "Need query parameters: region & serviceCode"
    ErrMsgAddPrice    = "Price matching that region, serviceCode, location & instanceType already exists"
)