package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// @Created 15/11/2021
// @Updated 
func GinHttpGetHandler(endpoint string, handler *gin.Context, queryKey, queryParam string) (*ResponseRecord, error) {
	// Setup gin in test mode, WIP: confirm if this statement is necessary
	gin.SetMode(gin.TestMode)

	// Create response recorder
	rr := httptest.NewRecorder()

	
	r := gin.New()

	// Parse api endpoint and http function handler from function arguments
	r.POST(endpoint, handler)

	// Create new http request
	httpRequest, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	if queryKey != "" && queryParam != "" {
		query := httpRequest.URL.Query()
		query.Add(queryKey, queryParam)
		httpRequest.URL.RawQuery = query.Encode()
	}

	r.ServeHTTP(rr, httpRequest)

	// Return http post response wapped in ResponseRecorder struct
	return rr, nil
}

// @Created 15/11/2021
// @Updated 
func GinHttpPostHandler(endpoint string, handler *gin.Context, requestBody []byte) (*ResponseRecord, error) {
	// Setup gin in test mode, WIP: confirm if this statement is necessary
	gin.SetMode(gin.TestMode)

	// Create response recorder
	rr := httptest.NewRecorder()

	r := gin.New()

	// Parse api endpoint and http function handler from function arguments
	r.POST(endpoint, handler)

	// Create new http request
	httpRequest, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	r.ServeHTTP(rr, httpRequest)

	// Return http post response wapped in ResponseRecorder struct
	return rr, nil
}