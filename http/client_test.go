package http

import (
	// "fmt"
	"testing"
	"time"

	"github.com/mushoffa/go-library/http"
)

var (
	client = http.NewHttpClient()
)

func TestHttpClientGet_Success(t *testing.T) {
	// client := http.NewHttpClient()
	response, err := client.Get("http://www.google.com", nil)
	if err != nil {
		t.Errorf("Error on http client GET request: %v", err)
	}

	if response == nil {
		t.Errorf("Error on http client GET response")
	}
}

func TestIsNetworkTimeout_Success(t *testing.T) {
	client.Timeout(5 * time.Second)
	_, err := client.Get("YOUT_TIMEOUT_API", nil)
	if err != nil {
		isTimeout := client.IsNetworkTimeout(err)
		if !isTimeout {
			t.Errorf("Supposed to be timeout")
		}
	}
}