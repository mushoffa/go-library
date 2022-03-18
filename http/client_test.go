package http

import (
	"testing"

	"github.com/mushoffa/go-library/http"
)

func TestHttpClientGet_Success(t *testing.T) {
	client := http.NewHttpClient()
	response, err := client.Get("http://www.google.com", nil)
	if err != nil {
		t.Errorf("Error on http client GET request: %v", err)
	}

	if response == nil {
		t.Errorf("Error on http client GET response")
	}
}