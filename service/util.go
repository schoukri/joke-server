package service

import (
	"net/http"
	"time"
)

const userAgent = "JokeServer/0.1"

// NewClient returns a new http.Client pre-configured with a short timeout for fetching requests from 3rd-party APIs.
func NewClient() *http.Client {
	return &http.Client{
		Timeout: 5 * time.Second,
	}
}

// addStandardHeaders adds standard headers to the specified Request, such as User Agent.
func addStandardHeaders(req *http.Request) {
	req.Header.Set("User-Agent", userAgent)
}

// RequestError is a custom error type which includes an http request status code
type RequestError struct {
	Err        error
	StatusCode int
}

// Error returns the string representation of the RequestError
func (r *RequestError) Error() string {
	return r.Err.Error()
}
