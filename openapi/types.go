package openapi

import (
	"etri5gc/fabric/common"
	"net/http"
	"net/url"
)


type ResponseHandlerFunc func(resp *http.Response, err error) (*Response, error)

type Request struct {
	Path         string
	Method       string
	Body         interface{}
	HeaderParams map[string]string
	QueryParams  url.Values
	FormParams   url.Values
	FormFileName string
	FileName     string
	FileBytes    []byte
	//a handler to handle a http response
	handler ResponseHandlerFunc
}

func (msg *Request) MsgType() int {
	return common.SERVICE_MSG_TYPE_OPENAPI
}

func DefaultRequest() *Request {
	ret := &Request{
		HeaderParams: make(map[string]string),
	}
	ret.HeaderParams["Content-Type"] = "application/json"
	ret.HeaderParams["Accept"] = "application/json"
	return ret
}

type Response struct {
	Raw        interface{}
	Body       []byte
	StatusCode int
}

func (msg *Response) MsgType() int {
	return common.SERVICE_MSG_TYPE_OPENAPI
}

type Consumer interface {
	Send(*Request) (*Response, error)
    DecodeBody(interface{}, []byte, string) error
}

