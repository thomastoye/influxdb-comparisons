package main

import "time"

var UseFastHttp = true

// HTTPClient is a reusable HTTP Client.
type HTTPClientCommon struct {
	Host       []byte
	HostString string
	debug      int
}

// HTTPClientDoOptions wraps options uses when calling `Do`.
type HTTPClientDoOptions struct {
	Debug                int
	PrettyPrintResponses bool
}

// HTTPClient interface.
type HTTPClient interface {
	HostString() string
	Ping()
	Do(q *Query, opts *HTTPClientDoOptions) (lag float64, err error)
}

func NewHTTPClient(host string, debug int, timeout time.Duration) HTTPClient {
	if UseFastHttp {
		return NewFastHTTPClient(host, debug, timeout)
	} else {
		return NewDefaultHTTPClient(host, debug, timeout)
	}
}

