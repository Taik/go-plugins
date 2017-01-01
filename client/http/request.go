package http

import (
	"github.com/micro/go-micro/client"
)

type httpRequest struct {
	service     string
	method      string
	contentType string
	request     interface{}
	opts        client.RequestOptions
}

func newHTTPRequest(service, method string, request interface{}, contentType string, reqOpts ...client.RequestOption) client.Request {
	var opts client.RequestOptions
	for _, o := range reqOpts {
		o(&opts)
	}

	return &httpRequest{
		service:     service,
		method:      method,
		request:     request,
		contentType: contentType,
		opts:        opts,
	}
}

func (h *httpRequest) ContentType() string {
	return h.contentType
}

func (h *httpRequest) Service() string {
	return h.service
}

func (h *httpRequest) Method() string {
	return h.method
}

func (h *httpRequest) Request() interface{} {
	return h.request
}

func (h *httpRequest) Stream() bool {
	return h.opts.Stream
}