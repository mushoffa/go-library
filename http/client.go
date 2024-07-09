package http

import (
	"crypto/tls"
	"time"
	"net"
	http1 "net/http"

	"github.com/parnurzeal/gorequest"
)

const (
	DefaultDebuggable 	= true
	
	// Unit in second(s)
	DefaultTimeout 		= 60

	// Unit in second(s)
	DefaultRetryBad 	= 1
)

type HttpClientService interface {
	Get(string, http1.Header) ([]byte, error)
	Post(string, interface{}, http1.Header) ([]byte, error)
	IsNetworkTimeout(error) bool
}

type HttpClient struct {
	isDebuggable 	bool
	timeout 		time.Duration
	retryBad 		int
}

func NewHttpClient() *HttpClient {
	return &HttpClient {
		isDebuggable: DefaultDebuggable,
		timeout 	: DefaultTimeout * time.Second,
		retryBad 	: DefaultRetryBad,
	}
}

func (c *HttpClient) WithDebugMode(isDebuggable bool) *HttpClient {
	c.isDebuggable = isDebuggable
	return c
}

func (c *HttpClient) Timeout(timeout time.Duration) *HttpClient {
	c.timeout = timeout
	return c
}

func (c *HttpClient) RetryBad(retry int) *HttpClient {
	c.retryBad = retry
	return c
}

func (c *HttpClient) Get(url string, headers http1.Header) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(c.isDebuggable)
	agent := request.Get(url)

	if headers != nil {
		//agent.Header = headers
	}

	_, body, errs := agent.
		Timeout(c.timeout).
		Retry(c.retryBad, time.Second, http1.StatusInternalServerError).
		End()

	if errs != nil {
		return nil, errs[0]
	}

	return []byte(body), nil
}

func (c *HttpClient) Post(url string, jsonData interface{}, headers http1.Header) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(c.isDebuggable)

	if url[:5] == "https" {
		request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}

	agent := request.Post(url)
	agent.Header.Add("Content-Type", "application/json")

	if headers != nil {
		//agent.Header = headers
	}

	_, body, errs := agent.
		Send(jsonData).
		Timeout(c.timeout).
		Retry(c.retryBad, time.Second, http1.StatusInternalServerError).
		End()

	if errs != nil {
		return nil, errs[0]
	}

	return []byte(body), nil
}

func (c *HttpClient) IsNetworkTimeout(err error) bool {
	if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
		return true
	}

	return false
}
