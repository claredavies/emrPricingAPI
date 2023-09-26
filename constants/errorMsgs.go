package constants

const (
    ErrMsgParamIDRequired    = "ID parameter is required."
    ErrMsgQueryIDRequired    = "Query ID parameter is required."
    ErrInvalidJSON = "Invalid JSON"
    ErrMsgQueryGetPrice    = "Need query parameters: region, serviceCode, location & instanceType"
    ErrMsgBodyAddPrice    = "Need body parameters: region, serviceCode, location & instanceType"
    ErrMsgQueryGetPrices    = "Need query parameters: region & serviceCode"
    ErrMsgAddPrice    = "Price matching that region, serviceCode, location & instanceType already exists"
)