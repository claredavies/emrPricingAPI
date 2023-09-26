package price

import (
    "encoding/json"
    "net/http/httptest"
    "testing"
    "net/url"
    "emrPricingAPI/models"
//     "fmt"
    "bytes"
    "github.com/labstack/echo/v4"
)

func getMockRequestResponseContext(method, path string) (*echo.Echo, *httptest.ResponseRecorder, echo.Context){
    e := echo.New()
    SetupRoutes(e)

    request := httptest.NewRequest(method, path, nil)
    response := httptest.NewRecorder()
    c := e.NewContext(request, response)

    return e, response, c
}

func getMockRequestResponseContextWithQueries(method, path string, queryParams map[string]string) (*echo.Echo, *httptest.ResponseRecorder, echo.Context) {
    e := echo.New()
    SetupRoutes(e)

    // Create a URL with query parameters
    u, err := url.Parse(path)
    if err != nil {
        panic(err) // Handle error appropriately
    }

    query := u.Query()
    for key, value := range queryParams {
        query.Add(key, value)
    }
    u.RawQuery = query.Encode()

    request := httptest.NewRequest(method, u.String(), nil)

    response := httptest.NewRecorder()
    c := e.NewContext(request, response)

    return e, response, c
}

func getMockRequestResponseContextWithRequestBody(method, path string, requestBody []byte) (*echo.Echo, *httptest.ResponseRecorder, echo.Context) {
    e := echo.New()
    SetupRoutes(e)

    // Create a request with the specified method, path, and request body
    request := httptest.NewRequest(method, path, bytes.NewReader(requestBody))
    request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON) // Set the content type to JSON

    response := httptest.NewRecorder()
    c := e.NewContext(request, response)

    return e, response, c
}

func getMockResponseError(t *testing.T, response *httptest.ResponseRecorder) map[string]string {
    var responseError map[string]string
    if err := json.Unmarshal(response.Body.Bytes(), &responseError); err != nil {
        t.Fatalf("Failed reading response: %s", err)
    }
    return responseError
}

func getMockResponsePrice(t *testing.T, response *httptest.ResponseRecorder) models.Price {
    var responsePrice models.Price
    if err := json.Unmarshal(response.Body.Bytes(), &responsePrice); err != nil {
        t.Fatalf("Failed reading response: %s", err)
    }
    return responsePrice
}

func getMockResponsePrices(t *testing.T, response *httptest.ResponseRecorder) []models.Price {
     var responsePrices []models.Price
     if err := json.Unmarshal(response.Body.Bytes(), &responsePrices); err != nil {
        t.Fatalf("Failed reading response: %s", err)
     }
     return responsePrices
}