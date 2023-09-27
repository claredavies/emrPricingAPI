package constants

import "errors"

var (
    ErrQueryParameterMissing = errors.New(ErrMsgQueryGetPrice)
    ErrNoMatchingResults     = errors.New(ErrMsgNoMatchingResultsGetPrice)
    ErrTooManyResultsReturned    = errors.New(ErrMsgTooManyResultsReturned)
)

const (
    ErrMsgPriceNotFound = "Price not found with that ID"
    ErrMsgParamIDRequired    = "ID parameter is required."
    ErrMsgQueryIDRequired    = "Query ID parameter is required."
    ErrInvalidJSON = "Invalid JSON"
    ErrMsgTooManyResultsReturned = "There should only be 1 price matching the query"
    ErrMsgQueryGetPrice    = "Need query parameters: serviceCode & instanceType"
    ErrMsgNoMatchingResultsGetPrice    = "No matching results for serviceCode & instanceType"
    ErrMsgQueryGetPrices    = "Need query parameters: serviceCode"
)