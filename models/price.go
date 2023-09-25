package models

type Price struct {
    ID              string  `json:"ID"`
    ServiceType     string  `json:"ServiceType"`
    InstanceType    string  `json:"InstanceType"`
    Market          string  `json:"Market"`
    Unit            string  `json:"Unit"`
    PricePerUnit    float64 `json:"PricePerUnit"`
    PriceDescription string  `json:"PriceDescription"`
    UpdatedAt       string  `json:"UpdatedAt"`
}

type AddPrice struct {
    Region          string  `json:"Region"`
    ServiceCode     string  `json:"ServiceCode"`
    Location        string  `json:"Location"`
    InstanceType    string  `json:"InstanceType"`
}
