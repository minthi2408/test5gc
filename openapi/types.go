package openapi

import (
    "net/http"
	"net/url"
    "etri5gc/fabric/common"
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
    handler      ResponseHandlerFunc
}

func (req *Request) Handler() ResponseHandlerFunc {
    return req.handler
}

func (msg *Request) MsgType() int {
    return common.SERVICE_MSG_TYPE_OPENAPI
}
type Response struct {
	Raw  interface{}
	Body []byte
	Code int
}

func (msg *Response) MsgType() int {
    return common.SERVICE_MSG_TYPE_OPENAPI
}

