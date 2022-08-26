package openapi

import (
	"etri5gc/fabric/common"
	"net/url"
)

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

	BodyBytes []byte      //encoded body
	Request   interface{} //dataplace protocol-bound request
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
	Response  interface{}
	BodyBytes []byte
}

func (msg *Response) MsgType() int {
	return common.SERVICE_MSG_TYPE_OPENAPI
}

type Consumer interface {
	Send(*Request) (*Response, error)
	DecodeBody(interface{}, []byte, string) error
}
